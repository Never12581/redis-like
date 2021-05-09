package analysis

type PingCmd struct {
	*Cmd
}

func (c *PingCmd) Deal() []byte {
	return Pong
}

func (c *PingCmd) paramInit() {
}
