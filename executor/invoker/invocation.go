package invoker

import (
	"context"
	"redis-like/executor/result"
)

type CallBackFunc func(ctx context.Context, invocation InvocationInter, inter result.ResultInter)

type InvocationInter interface {
	GetAttachments() map[string]interface{}
	GetAttachment(key string) interface{}
	GetAttachmentOrDefaultValue(key string, defaultValue interface{}) interface{}
	PutAttachment(key string, value interface{})
	AddCallbacks(backFunc CallBackFunc)
	OnFinished(ctx context.Context, inter result.ResultInter)
}

type Invocation struct {
	callbacks   chan CallBackFunc
	attachments map[string]interface{}
}

func NewInvocation() *Invocation {
	i := &Invocation{}
	i.attachments = make(map[string]interface{}, 0)
	i.callbacks = make(chan CallBackFunc)
	return i
}

func (ic *Invocation) OnFinished(ctx context.Context, inter result.ResultInter) {
	if len(ic.callbacks) == 0 {
		close(ic.callbacks)
		return
	}
	for {
		if callback, ok := <-ic.callbacks; ok {
			if callback != nil {
				callback(ctx, ic, inter)
			}
		} else {
			close(ic.callbacks)
			break
		}
	}
}

func (ic *Invocation) PutAttachment(key string, value interface{}) {
	ic.attachments[key] = value
}

func (ic *Invocation) GetAttachments() map[string]interface{} {
	return ic.attachments
}

func (ic *Invocation) GetAttachment(key string) interface{} {
	return ic.attachments[key]
}

func (ic *Invocation) GetAttachmentOrDefaultValue(key string, defaultValue interface{}) interface{} {
	value := ic.attachments[key]
	if value == nil {
		value = defaultValue
	}
	return value
}

func (ic *Invocation) AddCallbacks(callback CallBackFunc) {
	ic.callbacks <- callback
}
