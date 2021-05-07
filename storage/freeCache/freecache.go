package freeCache

import (
	"context"
	"github.com/coocood/freecache"
	"redis-like/storage"
)

const engine = "free_cache"

type FreeCache struct {
	cache *freecache.Cache
}

func NewFreeCache(cache *freecache.Cache) storage.Storage {
	return &FreeCache{cache: cache}
}

func (f *FreeCache) Del(context context.Context, key []byte) error {
	flag := f.cache.Del(key)
	if flag {
		return nil
	}
	return storage.NewError(storage.DelErrorText, engine, nil)
}

// fixme : get and set must be in a atom operator
func (f *FreeCache) Append(context context.Context, key []byte, value []byte) error {
	val, err := f.cache.Get(key)
	if err != nil {
		return storage.NewError(storage.AppendGetErrorText, engine, err)
	}
	val = append(val, value...)
	err = f.cache.Set(key, val, 0)
	if err != nil {
		return storage.NewError(storage.AppendSetErrorText, engine, err)
	}
	return nil
}

func (f *FreeCache) Get(context context.Context, key []byte) ([]byte, error) {
	val, err := f.cache.Get(key)
	if err != nil {
		return nil, storage.NewError(storage.GetErrorText, engine, err)
	}
	return val, nil
}

func (f *FreeCache) Set(context context.Context, key []byte, value []byte) error {
	err := f.cache.Set(key, value, 0)
	if err != nil {
		return storage.NewError(storage.SetErrorText, engine, err)
	}
	return nil
}
