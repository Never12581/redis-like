package protocol

import "context"

type AppendCmd struct {
	*Cmd
	key []byte
	val []byte
}

func (c *AppendCmd) Deal() []byte {
	err := c.storage.Append(context.Background(), c.key, c.val)

	if err == nil {
		out := OK
		return out
	} else {
		return CommonErr
	}
}

func (c *AppendCmd) paramInit() {
	c.key = c.Cmd.ParamBs[1]
	c.val = c.Cmd.ParamBs[2]
}
