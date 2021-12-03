package standalone_storage

import (
	"path/filepath"

	"github.com/Connor1996/badger"
	"github.com/pingcap-incubator/tinykv/kv/config"
	"github.com/pingcap-incubator/tinykv/kv/storage"
	"github.com/pingcap-incubator/tinykv/kv/util/engine_util"
	"github.com/pingcap-incubator/tinykv/proto/pkg/kvrpcpb"
)

//StandAloneStorage是单节点TinyKV实例的“存储”实现。事实并非如此与其他节点通信，所有数据都存储在本地。
type StandAloneStorage struct {
	conf 	*config.Config
	it 		*engine_util.BadgerIterator
}

type BadgerRead struct{
	txn *badger.Txn
	db  *badger.DB
	it	*engine_util.BadgerIterator
}

func NewStandAloneStorage(conf *config.Config) *StandAloneStorage {
	filepath.Join()
	return nil
}

func (s *StandAloneStorage) Start() error {
	// Your Code Here (1).
	return nil
} 
func (s *StandAloneStorage) Stop() error {
	// Your Code Here (1).
	return nil
}

func (s *StandAloneStorage) Reader(ctx *kvrpcpb.Context) (storage.StorageReader, error) {
	// Your Code Here (1).
	return nil, nil
}

func (s *StandAloneStorage) Write(ctx *kvrpcpb.Context, batch []storage.Modify) error {
	// Your Code Here (1).
	return nil
}
