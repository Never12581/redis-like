package storage

import (
	"fmt"
	"github.com/coocood/freecache"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"redis-like/config"
	"runtime/debug"
	"sync"
)

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
