package cmd

import (
	"context"
	"redis-like/constant"
	"redis-like/storage"
)

type GetCmd struct {
	key []byte
}

func (g *GetCmd) Init(bs [][]byte) error {
	g.key = bs[0]
	return nil
}

func (g *GetCmd) Deal(ctx context.Context) []byte {
	storage := storage.StorageInstance()
	bs, err := storage.Get(ctx, g.key)
	if err == nil {
		out := []byte("+")
		out = append(out, bs...)
		return out
	} else {
		return constant.NotFoundErr
	}
}
