package storage

import (
	"context"
	"github.com/syndtr/goleveldb/leveldb"
	"redis-like/config"
)

type LevelDB struct {
	db *leveldb.DB
}

func NewLevelDB(db *leveldb.DB) Storage {
	return &LevelDB{db: db}
}

func (l *LevelDB) Del(context context.Context, key []byte) error {
	err := l.db.Delete(key, nil)
	if err != nil {
		return newLevelDBError(DelErrorText, err)
	}
	return nil
}

func (l *LevelDB) Append(context context.Context, key []byte, value []byte) error {
	val, err := l.db.Get(key, nil)
	if err != nil {
		return newLevelDBError(AppendGetErrorText, err)
	}
	val = append(val, value...)
	err = l.db.Put(key, val, nil)
	if err != nil {
		return newLevelDBError(AppendSetErrorText, err)
	}
	return nil
}

func (l *LevelDB) Get(context context.Context, key []byte) ([]byte, error) {
	val, err := l.db.Get(key, nil)
	if err != nil {
		return nil, newLevelDBError(GetErrorText, err)
	}
	return val, nil
}

func (l *LevelDB) Set(context context.Context, key []byte, value []byte) error {
	err := l.db.Put(key, value, nil)
	if err != nil {
		return newLevelDBError(SetErrorText, err)
	}
	return nil
}

func newLevelDBError(text ErrorInfo, err error) error {
	return NewError(text, config.LevelDBEngine, err)
}
