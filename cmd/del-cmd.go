package cmd

import (
	"context"
	"redis-like/constant"
	"redis-like/storage"
)

type DelCmd struct {
	key []byte
}

func (d *DelCmd) Init(bs [][]byte) error {
	d.key = bs[0]
	return nil
}

func (d *DelCmd) Deal(ctx context.Context) []byte {
	storage := storage.StorageInstance()
	err := storage.Del(context.Background(), d.key)
	if err != nil {
		return constant.CommonErr
	} else {
		return constant.OK
	}
}
