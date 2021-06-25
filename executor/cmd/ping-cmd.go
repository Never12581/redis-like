package cmd

import (
	"context"
	"redis-like/executor"
)

type PingCmd struct {
}

func (p *PingCmd) Init(bs [][]byte) error {
	return nil
}

func (p *PingCmd) Deal(ctx context.Context) []byte {
	return executor.Pong
}
