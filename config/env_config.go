package config

import (
	"os"
	"strconv"
	"sync"
)

const (
	EnvEngineType    = "engine_type"
	EnvStoragePath   = "storage_path"
	EnvFreeCacheSize = "free_cache_size"
)

const (
	FreeCacheEngine = "free_cache"
	LevelDBEngine   = "levelDB"
)

const (
	defaultStoragePath   = "redis-like"
	defaultDBEngine      = FreeCacheEngine
	defaultFreeCacheSize = 100 * 1024 * 1024
)

var (
	envModeInstance *EnvMode
	once            sync.Once
)

func EnvConfigInstance() *EnvMode {
	if envModeInstance == nil {
		once.Do(func() {
			envModeInstance = new(EnvMode)

			engine, exist := os.LookupEnv(EnvEngineType)
			if !exist {
				engine = LevelDBEngine
			}
			envModeInstance.SetEngine(engine)

			storagePath, _ := os.LookupEnv(EnvStoragePath)
			envModeInstance.SetStoragePath(storagePath)

			freeCacheSize, _ := os.LookupEnv(EnvFreeCacheSize)
			envModeInstance.SetFreeCacheEngineSize(freeCacheSize)
		})
	}
	return envModeInstance
}

type EnvMode struct {
	engine              string
	storagePath         string
	freeCacheEngineSize int // freeCache 存储引擎的 内存型存储item 数量
}

func (m *EnvMode) Engine() (string, bool) {
	if len(m.engine) == 0 {
		return "", false
	}
	return m.engine, true
}

func (m *EnvMode) SetEngine(engine string) {
	if len(m.engine) == 0 {
		m.engine = defaultDBEngine
	} else {
		m.engine = engine
	}
}

func (m *EnvMode) StoragePath() (string, bool) {
	if len(m.storagePath) == 0 {
		return "", false
	}
	return m.storagePath, true
}

func (m *EnvMode) SetStoragePath(storagePath string) {
	if len(m.storagePath) == 0 {
		m.storagePath = defaultStoragePath
	} else {
		m.storagePath = storagePath
	}
}

func (m *EnvMode) FreeCacheEngineSize() (int, bool) {
	return m.freeCacheEngineSize, true
}

func (m *EnvMode) SetFreeCacheEngineSize(freeCacheEngineSize string) {

	size, err := strconv.Atoi(freeCacheEngineSize)
	if err != nil {
		m.freeCacheEngineSize = defaultFreeCacheSize
	} else {
		m.freeCacheEngineSize = size
	}
}
