package executor

import (
	"context"
	"log"
)

type InvokerInter interface {
	Invoke(ctx context.Context, invocation InvocationInter) ResultInter
	Callback() CallBackFunc
	hasNext() bool
	next() InvokerInter
}

// SimpleInvoker 简单demo
type SimpleInvoker struct {
	nextInvoker InvokerInter
	callback    CallBackFunc
}

func (i *SimpleInvoker) Invoke(ctx context.Context, invocation InvocationInter) ResultInter {
	log.Println("simple invoke start!")
	return nil
}

func (i *SimpleInvoker) Callback() CallBackFunc {
	return nil
}

func (i *SimpleInvoker) hasNext() bool {
	return i.nextInvoker != nil
}

func (i *SimpleInvoker) next() InvokerInter {
	return i.nextInvoker
}

// StorageInvoker 存储demo
type StorageInvoker struct {
	nextInvoker InvokerInter
}

func (s *StorageInvoker) Invoke(ctx context.Context, invocation InvocationInter) ResultInter {
	panic("implement me")
}

func (s *StorageInvoker) Callback() CallBackFunc {
	panic("implement me")
}

func (s *StorageInvoker) hasNext() bool {
	return s.nextInvoker != nil
}

func (s *StorageInvoker) next() InvokerInter {
	return s.nextInvoker
}
