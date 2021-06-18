package result

import "fmt"

const (
	NoResultError = "no result error."
)

type ResultInter interface {
	Success() bool
	Result() interface{}
	Error() error
	HasError() bool
}

type Result struct {
	success bool
	result  interface{}
	error   error
}

func (r *Result) Success() bool {
	return r.success
}

func (r *Result) Result() interface{} {
	return r.result
}

func (r *Result) Error() error {
	return r.error
}

func (r *Result) HasError() bool {
	return r.error == nil
}

// DefaultResult 默认返回
func DefaultResult() ResultInter {
	r := new(Result)
	r.success = false
	r.error = fmt.Errorf(NoResultError)
	return r
}

// ErrorResult 默认错误返回
func ErrorResult(err error) ResultInter {
	r := new(Result)
	r.success = false
	r.error = err
	return r
}

func SuccessResult(result interface{}) ResultInter {
	r := new(Result)
	r.success = true
	r.result = result
	return r
}
