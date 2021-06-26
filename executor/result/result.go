package result

import (
	"fmt"
	"redis-like/executor"
)

type ResultInter interface {
	Success() bool
	Result() []byte
	Error() error
	HasError() bool
}

type Result struct {
	success bool
	result  []byte
	error   error
}

func (r *Result) Success() bool {
	return r.success
}

func (r *Result) Result() []byte {
	return r.result
}

func (r *Result) Error() error {
	return r.error
}

func (r *Result) HasError() bool {
	return r.error == nil
}

func DefaultSuccessResult() ResultInter {
	r := new(Result)
	r.success = true
	return r
}

// DefaultResult 默认返回
func DefaultResult() ResultInter {
	r := new(Result)
	r.success = false
	r.error = fmt.Errorf(executor.NoResultError)
	return r
}

// ErrorResult 默认错误返回
func ErrorResult(err error) ResultInter {
	r := new(Result)
	r.success = false
	r.error = err
	return r
}

// SuccessResult 正确返回
func SuccessResult(result []byte) ResultInter {
	r := new(Result)
	r.success = true
	r.result = result
	return r
}

func SuccessWithoutResult() ResultInter {
	r := new(Result)
	r.success = true
	return r
}
