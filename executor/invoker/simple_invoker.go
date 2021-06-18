package invoker

import (
	"context"
	"log"
	"use-demo/frame/result"
)

// SimpleInvoker 简单demo
type SimpleInvoker struct {
	nextInvoker InvokerInter
	callback    CallBackFunc
}

func (i *SimpleInvoker) SetNext(inter InvokerInter) {
	i.nextInvoker = inter
}

func (i *SimpleInvoker) Invoke(ctx context.Context, invocation InvocationInter) result.ResultInter {
	log.Println("simple invoke start!")
	return result.DefaultResult()
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
