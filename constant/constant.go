package constant

import "errors"

const (
	NoResultError = "no result error."
)

var (
	RequestErr      = []byte("CommonErr request .")
	RequestStartErr = []byte("CommonErr request not start with * .")
	Pong            = []byte("Pong")
	OK              = []byte("OK")
	CommonErr       = []byte("cannot deal error .")
	UnsupportedErr  = []byte("the func unsupported .")
	NotFoundErr     = []byte("the key not found .")
)

var (
	UnsupportedCommandErr = errors.New("unsupported command error")
	ParamsGetError        = errors.New("source params get error!")
	ParamsAnalysisError   = errors.New("source params analysis error！")
)
