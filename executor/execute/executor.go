package execute

import (
	"context"
	"github.com/google/martian/log"
	"redis-like/executor/invoker"
	"redis-like/executor/result"
	"runtime/debug"
	"sync"
)

var (
	executor     Executor
	executorOnce sync.Once
)

func ExecutorInstance() Executor {
	executorOnce.Do(func() {
		executor = &SimpleExecutor{}
		inProtocolInvoker := invoker.InProtocolInvokerInstance()
		storageInvoker := invoker.StorageInvokerInstance()
		outProtocolInvoker := invoker.OutProtocolInvokerInstance()

		inProtocolInvoker.SetNext(storageInvoker)
		storageInvoker.SetNext(outProtocolInvoker)

		executor.SetInvoker(inProtocolInvoker)
	})

	return executor
}

type Executor interface {
	Execute(ctx context.Context, invocation invoker.InvocationInter) []byte
	SetInvoker(inter invoker.Invoker)
}

type SimpleExecutor struct {
	invoker invoker.Invoker
}

func (s *SimpleExecutor) SetInvoker(inter invoker.Invoker) {
	s.invoker = inter
}

func (s *SimpleExecutor) Execute(ctx context.Context, invocation invoker.InvocationInter) []byte {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf(string(debug.Stack()))
		}
	}()

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
		if !i.HasNext() {
			break
		}
		i = i.Next()
	}
	if r.Success() {
		invocation.OnFinished(ctx, r)
		return r.Result()[0]
	} else {
		log.Errorf("%v", r.Error())
		if r.Result() != nil && len(r.Result()) > 1 {
			return r.Result()[0]
		}
		return []byte{'-', 'u', 'n', 'e', 'x', 'p', 'e', 'c', 't', 'e', 'd', ' ', 'e', 'r', 'r', 'o', 'r', '.', '\r', '\n'}
	}
}
