package levelDB

import (
	"context"
	"github.com/syndtr/goleveldb/leveldb"
	"redis-like/storage"
)

const engine = "levelDB"

type LevelDB struct {
	db *leveldb.DB
}

func NewLevelDB(db *leveldb.DB) storage.Storage {
	return &LevelDB{db: db}
}

func (l *LevelDB) Del(context context.Context, key []byte) error {
	err := l.db.Delete(key, nil)
	if err != nil {
		return storage.NewError(storage.DelErrorText, engine, err)
	}
	return nil
}

func (l *LevelDB) Append(context context.Context, key []byte, value []byte) error {
	val, err := l.db.Get(key, nil)
	if err != nil {
		return storage.NewError(storage.AppendGetErrorText, engine, err)
	}
	val = append(val, value...)
	err = l.db.Put(key, val, nil)
	if err != nil {
		return storage.NewError(storage.AppendSetErrorText, engine, err)
	}
	return nil
}

func (l *LevelDB) Get(context context.Context, key []byte) ([]byte, error) {
	val, err := l.db.Get(key, nil)
	if err != nil {
		return nil, storage.NewError(storage.GetErrorText, engine, err)
	}
	return val, nil
}

func (l *LevelDB) Set(context context.Context, key []byte, value []byte) error {
	err := l.db.Put(key, value, nil)
	if err != nil {
		return storage.NewError(storage.SetErrorText, engine, err)
	}
	return nil
}
