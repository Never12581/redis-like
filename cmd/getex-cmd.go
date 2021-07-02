package cmd

import (
	"context"
	"redis-like/executor/result"
	"strconv"
	"strings"
)

type GetExCmd struct {
	key  []byte
	ex   bool
	exAt int64
	px   bool
	pxAt int64
}

func (g *GetExCmd) Init(bs [][]byte) error {
	g.key = bs[0]
	for i := 1; i < len(bs); i = i + 2 {
		temp := string(bs[i])
		ttl, err := strconv.ParseInt(string(bs[i+1]), 10, 64)
		if err != nil {
			return err
		}
		if strings.Compare(temp, "ex") == 0 {
			g.ex = true
			g.exAt = ttl
		} else if strings.Compare(temp, "px") == 0 {
			g.px = true
			g.pxAt = ttl
		}
	}
	return nil
}

func (g *GetExCmd) Deal(ctx context.Context) result.ResultInter {
	panic("implement me")
}
