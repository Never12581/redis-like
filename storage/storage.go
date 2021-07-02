package storage

import (
	"context"
	"fmt"
	"github.com/coocood/freecache"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"redis-like/config"
	"runtime/debug"
	"sync"
)

type ErrorInfo string

const (
	DelErrorText       = ErrorInfo("engine: %v , operator del failed , error info: %v .")
	AppendGetErrorText = ErrorInfo("engine: %v , operator append failed in get , error info: %v .")
	AppendSetErrorText = ErrorInfo("engine: %v , operator append failed in set , error info: %v .")
	GetErrorText       = ErrorInfo("engine: %v , operator get failed , error info: %v .")
	SetErrorText       = ErrorInfo("engine: %v , operator set failed , error info: %v .")
)

type Storage interface {
	Del(context context.Context, key []byte) error
	Append(context context.Context, key []byte, value []byte) error
	Get(context context.Context, key []byte) ([]byte, error)
	Set(context context.Context, key []byte, value []byte) error
	Close()
}

var (
	once    sync.Once
	storage Storage
)

func NewError(text ErrorInfo, engine string, err error) error {
	var errorInfo string
	if err != nil {
		errorInfo = err.Error()
	}
	return fmt.Errorf(string(text), engine, errorInfo)
}

func StorageInstance() Storage {
	once.Do(newStorage)
	return storage
}

func newStorage() {

	envMode := config.EnvConfigInstance()

	engine, _ := envMode.Engine()
	switch engine {
	case config.LevelDBEngine:

		storagePath, _ := envMode.StoragePath()
		db, err := leveldb.OpenFile(storagePath, nil)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		storage = NewLevelDB(db)

	case config.FreeCacheEngine:

		freeCacheSize, _ := envMode.FreeCacheEngineSize()
		cache := freecache.NewCache(freeCacheSize)
		debug.SetGCPercent(20)
		storage = NewFreeCache(cache)
	}

}
