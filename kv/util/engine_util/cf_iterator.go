package engine_util

import (
	"github.com/Connor1996/badger"
)

type CFItem struct {
	item      *badger.Item
	prefixLen int
}

//String返回项的字符串表示形式
func (i *CFItem) String() string {
	return i.item.String()
}

func (i *CFItem) Key() []byte {
	return i.item.Key()[i.prefixLen:]
}

func (i *CFItem) KeyCopy(dst []byte) []byte {
	return i.item.KeyCopy(dst)[i.prefixLen:]
}

func (i *CFItem) Version() uint64 {
	return i.item.Version()
}

func (i *CFItem) IsEmpty() bool {
	return i.item.IsEmpty()
}

func (i *CFItem) Value() ([]byte, error) {
	return i.item.Value()
}

func (i *CFItem) ValueSize() int {
	return i.item.ValueSize()
}

func (i *CFItem) ValueCopy(dst []byte) ([]byte, error) {
	return i.item.ValueCopy(dst)
}

func (i *CFItem) IsDeleted() bool {
	return i.item.IsDeleted()
}

func (i *CFItem) EstimatedSize() int64 {
	return i.item.EstimatedSize()
}

func (i *CFItem) UserMeta() []byte {
	return i.item.UserMeta()
}

type BadgerIterator struct {
	iter   *badger.Iterator
	prefix string
}

func NewCFIterator(cf string, txn *badger.Txn) *BadgerIterator {
	return &BadgerIterator{
		iter:   txn.NewIterator(badger.DefaultIteratorOptions),
		prefix: cf + "_",
	}
}

func (it *BadgerIterator) Item() DBItem {
	return &CFItem{
		item:      it.iter.Item(),
		prefixLen: len(it.prefix),
	}
}

func (it *BadgerIterator) Valid() bool { return it.iter.ValidForPrefix([]byte(it.prefix)) }

func (it *BadgerIterator) ValidForPrefix(prefix []byte) bool {
	return it.iter.ValidForPrefix(append([]byte(it.prefix), prefix...))
}

func (it *BadgerIterator) Close() {
	it.iter.Close()
}

func (it *BadgerIterator) Next() {
	it.iter.Next()
}

func (it *BadgerIterator) Seek(key []byte) {
	it.iter.Seek(append([]byte(it.prefix), key...))
}

func (it *BadgerIterator) Rewind() {
	it.iter.Rewind()
}

type DBIterator interface {
	//项返回指向当前键值对的指针。
	Item() DBItem
	//当迭代完成时，Valid返回false。
	Valid() bool
	//接下来将迭代器向前推进一步。总是在下一个（）之后检查它。Valid（）
	//以确保您有权访问有效的it.Item（）。
	Next()
	//Seek将查找提供的密钥（如果存在）。如果不存在，它将寻找下一个最小的键
	//大于所提供的。
	Seek([]byte)

	//关闭迭代器
	Close()
}

type DBItem interface {
	//键返回键。
	Key() []byte
	//KeyCopy返回项的键的副本，并将其写入dst切片。
	//如果传递了nil，或者dst的容量不足，则将分配并返回一个新的片。
	KeyCopy(dst []byte) []byte
	// Value retrieves the value of the item.
	Value() ([]byte, error)
	// ValueSize returns the size of the value.
	ValueSize() int
	// ValueCopy returns a copy of the value of the item from the value log, writing it to dst slice.
	// If nil is passed, or capacity of dst isn't sufficient, a new slice would be allocated and
	// returned.
	ValueCopy(dst []byte) ([]byte, error)
}
