package cmd

import (
	"context"
	"redis-like/executor/result"
)

type GetDelCmd struct {
	key []byte
}

func (g *GetDelCmd) Init(bs [][]byte) error {
	g.key = bs[0]
	return nil
}

func (g *GetDelCmd) Deal(ctx context.Context) result.ResultInter {
	panic("implement me")
}
