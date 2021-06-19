package invoker

import (
	"context"
	"redis-like/executor/result"
)

const (
	RequestParams  = "requestParams"
	AnalysisParams = "analysisParams"
	ExecuteMethod  = "executeMethod"
)

type InvokerInter interface {
	Invoke(ctx context.Context, invocation InvocationInter) result.ResultInter
	Callback() CallBackFunc
	SetNext(inter InvokerInter)
	HasNext() bool
	Next() InvokerInter
}
