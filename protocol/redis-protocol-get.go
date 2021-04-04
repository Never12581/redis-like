package protocol

import (
	"github.com/google/martian/log"
	levelErr "github.com/syndtr/goleveldb/leveldb/errors"
)

type GetCmd struct {
	*Cmd
	key []byte
}

func (c *GetCmd) Deal() []byte {
	val, err := c.db.Get(c.key, nil)

	if err == nil {
		out := []byte("+")
		out = append(out, val...)
		return out
	} else if err == levelErr.ErrNotFound {
		return NotFoundErr
	} else {
		log.Errorf(err.Error())
		return CommonErr
	}
}

func (c *GetCmd) paramInit() {
	if len(c.Cmd.ParamBs) == 2 {
		c.key = c.Cmd.ParamBs[1]
	}
}
