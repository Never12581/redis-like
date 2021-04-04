package protocol

type DelCmd struct {
	*Cmd
	key []byte
}

func (c *DelCmd) Deal() []byte {
	err := c.db.Delete(c.key, nil)
	if err != nil {
		return CommonErr
	} else {
		return OK
	}
}

func (c *DelCmd) paramInit() {
	c.key = c.Cmd.ParamBs[1]
}
