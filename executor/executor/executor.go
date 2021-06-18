package executor

import (
	"context"
	"use-demo/frame/invoker"
	"use-demo/frame/result"
)

type Executor interface {
	Execute(ctx context.Context, invocation invoker.InvocationInter) (interface{}, error)
	AddInvoker(inter invoker.InvokerInter)
}

type SimpleExecutor struct {
	invoker invoker.InvokerInter
}

func (s *SimpleExecutor) AddInvoker(inter invoker.InvokerInter) {
	s.invoker = inter
}

func (s *SimpleExecutor) Execute(ctx context.Context, invocation invoker.InvocationInter) result.ResultInter {
	i := s.invoker
	var r result.ResultInter
	for {
		r = i.Invoke(ctx, invocation)
		if !r.Success() || !i.HasNext() {
			break
		}
		invocation.AddCallbacks(i.Callback())
		i = i.Next()
	}
	if r.Success() {
		invocation.OnFinished(ctx, r)
	}
	return r
}
