package invoker

import (
	"context"
	"fmt"
	"redis-like/executor"
	"redis-like/executor/result"
	"strconv"
	"sync"
)

var (
	protocolInvoker *InProtocolInvoker
	protocolOnce    sync.Once
)

// InProtocolInvoker 协议处理invoker
type InProtocolInvoker struct {
	nextInvoker InvokerInter
}

func ProtocolInvokerInstance() *InProtocolInvoker {
	protocolOnce.Do(func() {
		protocolInvoker = &InProtocolInvoker{}
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
		bss, err := commonRespProtocolAnalysis(bs)
		if err == nil {
			invocation.PutAttachment(ExecuteMethod, string(bss[0]))
			invocation.PutAttachment(AnalysisParams, bss[1:])
			r = result.SuccessResult(nil)
		} else {
			r = result.ErrorResult(err)
		}
	} else {
		r = result.ErrorResult(fmt.Errorf(executor.ParamsGetError))
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

// 通用解析 ---> 解析为 [][]byte
func commonRespProtocolAnalysis(bs []byte) ([][]byte, error) {
	bsLength := len(bs)
	firstR := findFirstR(bs, 0, bsLength)
	if firstR == -1 {
		return nil, fmt.Errorf(executor.ParamsAnalysisError)
	}
	// 本次数组中参数长度
	paramLength, err := strconv.Atoi(string(bs[1:firstR]))
	if err != nil {
		return nil, err
	}
	// 逐个解析内容
	paramContents := make([][]byte, paramLength)
	tempStart := firstR + 2 // the first position of /n
	tempEnd := bsLength
	for i := 0; i < paramLength; i++ {
		bbs, nextOffset := analysisParamAndNextOffset(bs, tempStart, tempEnd)
		if nextOffset == -1 {
			return nil, fmt.Errorf(executor.ParamsAnalysisError)
		}
		tempStart = nextOffset
		paramContents[i] = bbs
	}
	return paramContents, nil
}

func findFirstR(bs []byte, start, end int) int {
	return findByteIndex(bs, start, end, '*', '\r')
}

// 返回当前次的 []byte 内容 ， 和对应的下一个start的偏移量
func analysisParamAndNextOffset(bs []byte, start, end int) ([]byte, int) {
	index := findByteIndex(bs, start, end, '$', '\r')
	if index == -1 {
		return nil, index
	}
	length, err := strconv.Atoi(string(bs[start+1 : index]))
	if err != nil {
		return nil, -1
	}
	bbs := bs[index+2 : index+2+length]
	offset := index + 2 + length + 2
	return bbs, offset
}

func findByteIndex(bs []byte, start, end int, startByte, endByte byte) int {
	if bs[start] != startByte {
		return -1
	}
	for ind, val := range bs[start:end] {
		if val == endByte {
			return ind + start
		}
	}
	return -1
}
