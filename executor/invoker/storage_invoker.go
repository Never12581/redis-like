package invoker

import (
	"context"
	"redis-like/cmd"
	"redis-like/executor/result"
	"sync"
)

var (
	storageInvoker *StorageInvoker
	storageOnce    sync.Once
)

// StorageInvoker 存储demo
type StorageInvoker struct {
	nextInvoker Invoker
}

func StorageInvokerInstance() *StorageInvoker {
	storageOnce.Do(func() {
		storageInvoker = &StorageInvoker{}
	})
	return storageInvoker
}

func (s *StorageInvoker) SetNext(inter Invoker) {
	s.nextInvoker = inter
}

func (s *StorageInvoker) Invoke(ctx context.Context, invocation InvocationInter) result.ResultInter {
	cmd := invocation.GetAttachment(ExecuteCmd).(cmd.Cmd)
	r := cmd.Deal(ctx)
	invocation.PutAttachment(SourceResult, r)
	return r
}

func (s *StorageInvoker) Callback() CallBackFunc {
	return nil
}

func (s *StorageInvoker) HasNext() bool {
	return s.nextInvoker != nil
}

func (s *StorageInvoker) Next() Invoker {
	return s.nextInvoker
}
