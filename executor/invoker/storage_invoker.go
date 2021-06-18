package invoker

import (
	"context"
	"use-demo/frame/cmd"
	"use-demo/frame/result"
)

// StorageInvoker 存储demo
type StorageInvoker struct {
	nextInvoker InvokerInter
}

func (s *StorageInvoker) SetNext(inter InvokerInter) {
	s.nextInvoker = inter
}

func (s *StorageInvoker) Invoke(ctx context.Context, invocation InvocationInter) result.ResultInter {
	//executeMethod := invocation.GetAttachment(ExecuteMethod).(string)
	//analysisParams := invocation.GetAttachment(AnalysisParams).([][]byte)

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
