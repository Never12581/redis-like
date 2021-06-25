package executor

import "errors"

const (
	NoResultError = "no result error."

	ParamsGetError      = "source params get error!"
	ParamsAnalysisError = "source params analysis error！"
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
	UnsupportedCommandErr = errors.New("unsupported command error")
)
