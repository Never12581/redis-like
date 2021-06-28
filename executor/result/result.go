package result

type ResultInter interface {
	Success() bool
	Result() [][]byte
	Error() error
}

type Result struct {
	success bool
	result  [][]byte
	error   error
}

func (r *Result) Success() bool {
	return r.success
}

func (r *Result) Result() [][]byte {
	return r.result
}

func (r *Result) Error() error {
	return r.error
}

// ErrorResult 默认错误返回
func ErrorResult(err error) ResultInter {
	if err != nil {
		r := new(Result)
		r.success = false
		r.error = err
		return r
	}
	return SuccessWithoutResult()
}

// SuccessResult 正确返回
func SuccessResult(result [][]byte) ResultInter {
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

// SuccessAndErrorResult 将error赋值，success为true
func SuccessAndErrorResult(result [][]byte, err error) ResultInter {
	r := new(Result)
	r.success = true
	r.result = result
	return r
}
