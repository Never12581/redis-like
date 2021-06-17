package executor

type ResultInter interface {
	Success() bool
	Result() interface{}
	Error() error
	HasError() bool
}
