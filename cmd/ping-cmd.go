package cmd

import (
	"context"
	"redis-like/constant"
)

type PingCmd struct {
}

func (p *PingCmd) Init(bs [][]byte) error {
	return nil
}

func (p *PingCmd) Deal(ctx context.Context) []byte {
	return constant.Pong
}
