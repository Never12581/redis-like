package protocol

import "context"

type DelCmd struct {
	*Cmd
	key []byte
}

func (c *DelCmd) Deal() []byte {
	err := c.storage.Del(context.Background(), c.key)
	if err != nil {
		return CommonErr
	} else {
		return OK
	}
}

func (c *DelCmd) paramInit() {
	c.key = c.Cmd.ParamBs[1]
}
