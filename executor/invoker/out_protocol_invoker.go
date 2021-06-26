package invoker

import (
	"context"
	"redis-like/executor/result"
)

type OutProtocolInvoker struct {
	nextInvoker InvokerInter
}

func (o *OutProtocolInvoker) Invoke(ctx context.Context, invocation InvocationInter) result.ResultInter {
	return nil
}

func (o *OutProtocolInvoker) Callback() CallBackFunc {
	return nil
}

func (o *OutProtocolInvoker) SetNext(inter InvokerInter) {
	o.nextInvoker = inter
}

func (o *OutProtocolInvoker) HasNext() bool {
	return o.nextInvoker == nil
}

func (o *OutProtocolInvoker) Next() InvokerInter {
	return o.nextInvoker
}
