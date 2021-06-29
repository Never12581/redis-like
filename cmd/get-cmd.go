package cmd

import (
	"context"
	"redis-like/executor/result"
	"redis-like/storage"
)

type GetCmd struct {
	key []byte
}

func (g *GetCmd) Init(bs [][]byte) error {
	g.key = bs[0]
	return nil
}

func (g *GetCmd) Deal(ctx context.Context) result.ResultInter {
	storage := storage.StorageInstance()
	bs, err := storage.Get(ctx, g.key)
	var r result.ResultInter
	if err == nil {
		bss := make([][]byte, 0)
		bss = append(bss, bs)
		r = result.SuccessResult(bss)
	} else {
		r = result.ErrorResult(err)
	}
	return r
}
