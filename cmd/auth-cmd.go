package cmd

import (
	"context"
	"redis-like/executor/result"
)

type AuthCmd struct {
	username []byte
	password []byte
}

func (a *AuthCmd) Init(bs [][]byte) error {
	if len(bs) == 2 {
		a.username = bs[0]
		a.password = bs[1]
	} else {
		a.password = bs[0]
	}
	return nil
}

func (a *AuthCmd) Deal(ctx context.Context) result.ResultInter {
	panic("implement auth-cmd")
}
