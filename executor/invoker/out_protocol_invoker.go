package invoker

import (
	"context"
	"redis-like/executor/result"
	"redis-like/protocol"
	"sync"
)

var (
	outProtocolInvoker *OutProtocolInvoker
	outProtocolOnce    sync.Once
)

type OutProtocolInvoker struct {
	nextInvoker InvokerInter
	resp        *protocol.RespProtocol
}

func OutProtocolInvokerInstance() *OutProtocolInvoker {
	outProtocolOnce.Do(func() {
		outProtocolInvoker = &OutProtocolInvoker{
			resp: protocol.RespProtocolInstance(),
		}
	})
	return outProtocolInvoker
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
