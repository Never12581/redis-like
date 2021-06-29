package cmd

import (
	"context"
	"redis-like/constant"
	"redis-like/executor/result"
	"redis-like/storage"
)

type DelCmd struct {
	key []byte
}

func (d *DelCmd) Init(bs [][]byte) error {
	d.key = bs[0]
	return nil
}

func (d *DelCmd) Deal(ctx context.Context) result.ResultInter {
	storage := storage.StorageInstance()
	err := storage.Del(context.Background(), d.key)
	var r result.ResultInter
	if err == nil {
		bss := make([][]byte, 0)
		bss = append(bss, constant.OK)
		r = result.SuccessResult(bss)
	} else {
		r = result.ErrorResult(err)
	}
	return r
}
