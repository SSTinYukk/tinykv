package engine_util

import (
	"os"

	"github.com/Connor1996/badger"
	"github.com/pingcap-incubator/tinykv/log"
)

//Engines保留unistore使用的引擎的引用和数据。
//所有引擎都是badger键/值数据库。
//路径字段是存储数据的文件系统路径。
type Engines struct {
	//数据，包括已提交（即跨其他节点提交）和未提交（即仅存在）的数据
	//本地）。
	Kv     *badger.DB
	KvPath string
	//Raft使用的元数据。
	Raft     *badger.DB
	RaftPath string
}

func NewEngines(kvEngine, raftEngine *badger.DB, kvPath, raftPath string) *Engines {
	return &Engines{
		Kv:       kvEngine,
		KvPath:   kvPath,
		Raft:     raftEngine,
		RaftPath: raftPath,
	}
}

func (en *Engines) WriteKV(wb *WriteBatch) error {
	return wb.WriteToDB(en.Kv)
}

func (en *Engines) WriteRaft(wb *WriteBatch) error {
	return wb.WriteToDB(en.Raft)
}

func (en *Engines) Close() error {
	if err := en.Kv.Close(); err != nil {
		return err
	}
	if err := en.Raft.Close(); err != nil {
		return err
	}
	return nil
}

func (en *Engines) Destroy() error {
	if err := en.Close(); err != nil {
		return err
	}
	if err := os.RemoveAll(en.KvPath); err != nil {
		return err
	}
	if err := os.RemoveAll(en.RaftPath); err != nil {
		return err
	}
	return nil
}

//CreateDB在子路径的磁盘上创建一个新的Badger DB。
func CreateDB(path string, raft bool) *badger.DB {
	opts := badger.DefaultOptions
	if raft {
		//不需要为raft引擎编写blob，因为它将很快被删除。
		opts.ValueThreshold = 0
	}
	opts.Dir = path
	opts.ValueDir = opts.Dir
	if err := os.MkdirAll(opts.Dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
