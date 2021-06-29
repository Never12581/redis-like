package invoker

import (
	"context"
	"redis-like/constant"
	"redis-like/executor/protocol"
	"redis-like/executor/result"
	"sync"
)

var (
	protocolInvoker *InProtocolInvoker
	protocolOnce    sync.Once
)

// InProtocolInvoker 协议处理invoker
type InProtocolInvoker struct {
	nextInvoker InvokerInter
	resp        *protocol.RespProtocol
}

func InProtocolInvokerInstance() *InProtocolInvoker {
	protocolOnce.Do(func() {
		protocolInvoker = &InProtocolInvoker{
			resp: protocol.RespProtocolInstance(),
		}
	})
	return protocolInvoker
}

func (p *InProtocolInvoker) SetNext(inter InvokerInter) {
	p.nextInvoker = inter
}

func (p *InProtocolInvoker) Invoke(ctx context.Context, invocation InvocationInter) result.ResultInter {
	bs, ok := invocation.GetAttachment(RequestParams).([]byte)
	var r result.ResultInter
	if ok {
		c, err := p.resp.UnPacket(bs)
		if err != nil {
			r = result.ErrorResult(err)
		} else {
			invocation.PutAttachment(ExecuteCmd, c)
			r = result.SuccessWithoutResult()
		}
	} else {
		r = result.ErrorResult(constant.ParamsGetError)
	}
	return r
}

func (p *InProtocolInvoker) Callback() CallBackFunc {
	return nil
}

func (s *InProtocolInvoker) HasNext() bool {
	return s.nextInvoker != nil
}

func (s *InProtocolInvoker) Next() InvokerInter {
	return s.nextInvoker
}
