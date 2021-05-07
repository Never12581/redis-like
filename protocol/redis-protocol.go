package protocol

import (
	"redis-like/storage"
)

type CmdDeal struct {
	cf CmdFunc
}

func (cd CmdDeal) Deal(db storage.Storage) []byte {
	cd.cf.setDB(db)
	cd.cf.paramInit()
	return cd.cf.Deal()
}

type CmdFunc interface {
	Deal() []byte
	setDB(storage.Storage)
	paramInit()
}

// RESP 协议解析
// default deal
type Cmd struct {
	ParamBs [][]byte // 参数值 如：ex 200
	storage storage.Storage
}

func (c *Cmd) Deal() []byte {
	return UnsupportedErr
}

func (c *Cmd) setDB(db storage.Storage) {
	c.storage = db
}

func (c *Cmd) paramInit() {
}
