package cmd

import (
	"context"
	"errors"
)

var (
	RequestErr      = []byte("-CommonErr request .")
	RequestStartErr = []byte("-CommonErr request not start with * .")
	Pong            = []byte("+Pong")
	OK              = []byte("+OK")
	CommonErr       = []byte("-cannot deal error .")
	UnsupportedErr  = []byte("-the func unsupported .")
	NotFoundErr     = []byte("-the key not found .")
)

var (
	UnsupportedCommandErr = errors.New("unsupported commond error")
)

type Cmd interface {
	Init(bs [][]byte) error
	Deal(ctx context.Context) []byte
}

type initFunc func() Cmd

var routeInfo map[string]initFunc

func init() {
	routeInfo = make(map[string]initFunc)
	routeInfo["ping"] = pingCmdInit
	routeInfo["get"] = getCmdInit
	routeInfo["set"] = setCmdInit
	routeInfo["append"] = appendCmdInit
	routeInfo["del"] = delCmdInit
}

func delCmdInit() Cmd {
	return new(DelCmd)
}

func appendCmdInit() Cmd {
	return new(AppendCmd)
}

func setCmdInit() Cmd {
	return new(SetCmd)
}

func getCmdInit() Cmd {
	return new(GetCmd)
}

func pingCmdInit() Cmd {
	return new(PingCmd)
}

func GeneratorCmd(executeMethod string, analysisParams [][]byte) (Cmd, error) {
	fn := routeInfo[executeMethod]
	if fn == nil {
		return nil, UnsupportedCommandErr
	}
	c := fn()
	err := c.Init(analysisParams)
	return c, err
}
