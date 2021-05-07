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
		return newError(storage.DelErrorText, err)
	}
	return nil
}

func (l *LevelDB) Append(context context.Context, key []byte, value []byte) error {
	val, err := l.db.Get(key, nil)
	if err != nil {
		return newError(storage.AppendGetErrorText, err)
	}
	val = append(val, value...)
	err = l.db.Put(key, val, nil)
	if err != nil {
		return newError(storage.AppendSetErrorText, err)
	}
	return nil
}

func (l *LevelDB) Get(context context.Context, key []byte) ([]byte, error) {
	val, err := l.db.Get(key, nil)
	if err != nil {
		return nil, newError(storage.GetErrorText, err)
	}
	return val, nil
}

func (l *LevelDB) Set(context context.Context, key []byte, value []byte) error {
	err := l.db.Put(key, value, nil)
	if err != nil {
		return newError(storage.SetErrorText, err)
	}
	return nil
}

func newError(text storage.ErrorInfo, err error) error {
	return storage.NewError(text, engine, err)
}
