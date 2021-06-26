package invoker

import (
	"context"
	"redis-like/executor/result"
)

const (
	// 原始参数
	RequestParams = "requestParams"
	// 解析后参数
	AnalysisParams = "analysisParams"
	// 执行方法
	ExecuteMethod = "executeMethod"
	// 原始返回值
	SourceResult = "sourceResult"
)

type InvokerInter interface {
	Invoke(ctx context.Context, invocation InvocationInter) result.ResultInter
	Callback() CallBackFunc
	SetNext(inter InvokerInter)
	HasNext() bool
	Next() InvokerInter
}
