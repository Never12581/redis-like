package invoker

import (
	"context"
	"log"
	"redis-like/executor/result"
	"sync"
)

var (
	simpleInvoker *SimpleInvoker
	simpleOnce    sync.Once
)

// SimpleInvoker 简单demo
type SimpleInvoker struct {
	nextInvoker InvokerInter
}

func SimpleInvokerInstance() *SimpleInvoker {
	simpleOnce.Do(func() {
		simpleInvoker = &SimpleInvoker{}
	})
	return simpleInvoker
}

func (i *SimpleInvoker) SetNext(inter InvokerInter) {
	i.nextInvoker = inter
}

func (i *SimpleInvoker) Invoke(ctx context.Context, invocation InvocationInter) result.ResultInter {
	log.Println("simple invoke start!")
	return result.DefaultSuccessResult()
}

func (i *SimpleInvoker) Callback() CallBackFunc {
	return nil
}

func (i *SimpleInvoker) HasNext() bool {
	return i.nextInvoker != nil
}

func (i *SimpleInvoker) Next() InvokerInter {
	return i.nextInvoker
}
