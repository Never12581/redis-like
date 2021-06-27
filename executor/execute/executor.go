package execute

import (
	"context"
	"github.com/google/martian/log"
	"redis-like/constant"
	"redis-like/executor/invoker"
	"redis-like/executor/result"

	"sync"
)

var (
	executor     Executor
	executorOnce sync.Once
)

func ExecutorInstance() Executor {
	executorOnce.Do(func() {
		executor = &SimpleExecutor{}
		protocolInvoker := invoker.ProtocolInvokerInstance()
		storageInvoker := invoker.StorageInvokerInstance()

		protocolInvoker.SetNext(storageInvoker)
		executor.SetInvoker(protocolInvoker)
	})

	return executor
}

type Executor interface {
	Execute(ctx context.Context, invocation invoker.InvocationInter) []byte
	SetInvoker(inter invoker.InvokerInter)
}

type SimpleExecutor struct {
	invoker invoker.InvokerInter
}

func (s *SimpleExecutor) SetInvoker(inter invoker.InvokerInter) {
	s.invoker = inter
}

func (s *SimpleExecutor) Execute(ctx context.Context, invocation invoker.InvocationInter) []byte {
	i := s.invoker
	var r result.ResultInter
	for {
		r = i.Invoke(ctx, invocation)
		if !r.Success() || !i.HasNext() {
			break
		}
		if i.Callback() != nil {
			invocation.AddCallbacks(i.Callback())
		}
		i = i.Next()
	}
	if r.Success() {
		invocation.OnFinished(ctx, r)
		return r.Result()
	}
	log.Errorf("%v", r.Error())
	return constant.CommonErr
}
