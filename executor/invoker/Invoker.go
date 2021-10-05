package invoker

import (
	"context"
	"redis-like/executor/result"
)

const (
	// RequestParams 原始参数
	RequestParams = "requestParams"
	// AnalysisParams 解析后参数
	AnalysisParams = "analysisParams"
	// ExecuteMethod 执行方法
	ExecuteMethod = "executeMethod"
	// SourceResult 原始返回值
	SourceResult = "sourceResult"
	// ExecuteCmd 执行方法对象
	ExecuteCmd = "executeCmd"
)

//Invoker 执行核心类
type Invoker interface {
	Invoke(ctx context.Context, invocation InvocationInter) result.ResultInter
	Callback() CallBackFunc
	SetNext(inter Invoker)
	HasNext() bool
	Next() Invoker
}
