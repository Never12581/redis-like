package cmd

import (
	"context"
	"redis-like/storage"
)

type AppendCmd struct {
	key []byte
	val []byte
}

func (a *AppendCmd) Init(bs [][]byte) error {
	a.key = bs[0]
	a.val = bs[1]
	return nil
}

func (a *AppendCmd) Deal(ctx context.Context) []byte {
	storage := storage.StorageInstance()
	err := storage.Append(context.Background(), a.key, a.val)
	if err == nil {
		out := OK
		return out
	} else {
		return CommonErr
	}
}
