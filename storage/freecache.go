package storage

import (
	"context"
	"github.com/coocood/freecache"
	"redis-like/config"
)

type FreeCache struct {
	cache *freecache.Cache
}

func NewFreeCache(cache *freecache.Cache) Storage {
	return &FreeCache{cache: cache}
}

func (f *FreeCache) Del(context context.Context, key []byte) error {
	flag := f.cache.Del(key)
	if flag {
		return nil
	}
	return newFreeCacheError(DelErrorText, nil)
}

// fixme : get and set must be in a atom operator
func (f *FreeCache) Append(context context.Context, key []byte, value []byte) error {
	val, err := f.cache.Get(key)
	if err != nil {
		return newFreeCacheError(AppendGetErrorText, err)
	}
	val = append(val, value...)
	err = f.cache.Set(key, val, 0)
	if err != nil {
		return newFreeCacheError(AppendSetErrorText, err)
	}
	return nil
}

func (f *FreeCache) Get(context context.Context, key []byte) ([]byte, error) {
	val, err := f.cache.Get(key)
	if err != nil {
		return nil, newFreeCacheError(GetErrorText, err)
	}
	return val, nil
}

func (f *FreeCache) Set(context context.Context, key []byte, value []byte) error {
	err := f.cache.Set(key, value, 0)
	if err != nil {
		return newFreeCacheError(SetErrorText, err)
	}
	return nil
}

func newFreeCacheError(text ErrorInfo, err error) error {
	return NewError(text, config.FreeCacheEngine, err)
}
