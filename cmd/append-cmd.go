package cmd

import (
	"context"
	"redis-like/constant"
	"redis-like/executor/result"
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

func (a *AppendCmd) Deal(ctx context.Context) result.ResultInter {
	storage := storage.StorageInstance()
	err := storage.Append(context.Background(), a.key, a.val)
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
