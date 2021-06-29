package cmd

import (
	"context"
	"redis-like/constant"
	"redis-like/executor/result"
)

type PingCmd struct {
}

func (p *PingCmd) Init(bs [][]byte) error {
	return nil
}

func (p *PingCmd) Deal(ctx context.Context) result.ResultInter {
	bss := make([][]byte, 0)
	bss = append(bss, constant.Pong)
	return result.SuccessResult(bss)
}
