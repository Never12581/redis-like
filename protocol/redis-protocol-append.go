package protocol

type AppendCmd struct {
	*Cmd
	key []byte
	val []byte
}

func (c *AppendCmd) Deal() []byte {
	sourceVal, err := c.db.Get(c.key, nil)
	if err != nil {
		return CommonErr
	}
	err = c.db.Put(c.key, append(sourceVal, c.val...), nil)

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
