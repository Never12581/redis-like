package analysis

import (
	"context"
	"strconv"
)

// EX seconds-设置指定的到期时间，以秒为单位。
// PX 毫秒-设置指定的到期时间（以毫秒为单位）。
// NX -仅设置不存在的密钥。
// XX -仅设置已存在的密钥。
type SetCmd struct {
	*Cmd
	key   []byte
	value []byte
	ex    bool
	exTtl int64
	px    bool
	pxTtl int64
	nx    bool
	xx    bool
}

func (c *SetCmd) Deal() []byte {
	err := c.storage.Set(context.Background(), c.key, c.value)
	if err == nil {
		return OK
	} else {
		return CommonErr
	}
}

// 到这里，已经肯定是 set 操作，可以使用一些硬编码操作
func (c *SetCmd) paramInit() {
	c.key = c.Cmd.ParamBs[1]
	if len(c.Cmd.ParamBs) >= 3 {
		c.value = c.Cmd.ParamBs[2]
	}
	for i := 3; i < len(c.Cmd.ParamBs); i = i + 2 {
		temp := string(c.Cmd.ParamBs[i])
		if temp == "ex" {
			c.ex = true
			c.exTtl, _ = strconv.ParseInt(string(c.Cmd.ParamBs[i+1]), 10, 64)
		} else if temp == "px" {
			c.px = true
			c.pxTtl, _ = strconv.ParseInt(string(c.Cmd.ParamBs[i+1]), 10, 64)
		} else if temp == "nx" {
			c.nx = true
		} else if temp == "xx" {
			c.xx = true
		}
	}
}
