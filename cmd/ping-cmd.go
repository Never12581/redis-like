package cmd

import (
	"context"
	"redis-like/constant"
	"redis-like/executor/result"
)

type PingCmd struct {
	value []byte
}

func (p *PingCmd) Init(bs [][]byte) error {
	if bs != nil && len(bs) >= 1 {
		p.value = bs[0]
	}
	return nil
}

func (p *PingCmd) Deal(ctx context.Context) result.ResultInter {
	bss := make([][]byte, 0)
	if p.value == nil {
		bss = append(bss, constant.Pong)
	} else {
		bss = append(bss, p.value)
	}
	return result.SuccessResult(bss)
}
