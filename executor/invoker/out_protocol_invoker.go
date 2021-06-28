package invoker

import (
	"context"
	"redis-like/executor/protocol"
	"redis-like/executor/result"
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
	resultInter := invocation.GetAttachment(SourceResult).(result.ResultInter)
	return o.resp.Packet(resultInter)
}

func (o *OutProtocolInvoker) Callback() CallBackFunc {
	return nil
}

func (o *OutProtocolInvoker) SetNext(inter InvokerInter) {
	o.nextInvoker = inter
}

func (o *OutProtocolInvoker) HasNext() bool {
	return o.nextInvoker != nil
}

func (o *OutProtocolInvoker) Next() InvokerInter {
	return o.nextInvoker
}
