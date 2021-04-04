package protocol

import "github.com/syndtr/goleveldb/leveldb"

type CmdDeal struct {
	cf CmdFunc
}

func (cd CmdDeal) Deal(db *leveldb.DB) []byte {
	cd.cf.setDB(db)
	cd.cf.paramInit()
	return cd.cf.Deal()
}

type CmdFunc interface {
	Deal() []byte
	setDB(*leveldb.DB)
	paramInit()
}

// RESP 协议解析
// default deal
type Cmd struct {
	ParamBs [][]byte // 参数值 如：ex 200
	db      *leveldb.DB
}

func (c *Cmd) Deal() []byte {
	return UnsupportedErr
}

func (c *Cmd) setDB(db *leveldb.DB) {
	c.db = db
}

func (c *Cmd) paramInit() {
}
