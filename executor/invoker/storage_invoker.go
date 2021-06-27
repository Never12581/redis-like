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
	nextInvoker InvokerInter
}

func StorageInvokerInstance() *StorageInvoker {
	storageOnce.Do(func() {
		storageInvoker = &StorageInvoker{}
	})
	return storageInvoker
}

func (s *StorageInvoker) SetNext(inter InvokerInter) {
	s.nextInvoker = inter
}

func (s *StorageInvoker) Invoke(ctx context.Context, invocation InvocationInter) result.ResultInter {
	cmd := invocation.GetAttachment(ExecuteCmd).(cmd.Cmd)
	bs := cmd.Deal(ctx)
	invocation.PutAttachment(SourceResult, bs)
	return result.SuccessResult(bs)
}

func (s *StorageInvoker) Callback() CallBackFunc {
	return nil
}

func (s *StorageInvoker) HasNext() bool {
	return s.nextInvoker != nil
}

func (s *StorageInvoker) Next() InvokerInter {
	return s.nextInvoker
}
