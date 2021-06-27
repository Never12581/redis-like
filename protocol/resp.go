package protocol

import (
	"redis-like/cmd"
	"redis-like/constant"
	"redis-like/executor/result"
	"strconv"
	"sync"
)

var (
	instance         *RespProtocol
	respProtocolOnce sync.Once
)

type RespProtocol struct {
}

func RespProtocolInstance() *RespProtocol {
	respProtocolOnce.Do(func() {
		instance = &RespProtocol{}
	})
	return instance
}

// *3\r\n$3\r\nset\r\n$1\r\na\r\n$6\r\nbcdefg
func (r *RespProtocol) UnPacket(bs []byte) (cmd.Cmd, error) {
	// 字节长度
	bsLength := len(bs)
	if bsLength == 0 || bs[0] != '*' {
		return nil, constant.ParamsAnalysisError
	}
	bss, err := commonRespProtocolAnalysis(bs)
	if err != nil {
		return nil, err
	}
	executeMethod := string(bss[0])
	analysisParams := bss[1:]
	return cmd.GeneratorCmd(executeMethod, analysisParams)
}

func (r *RespProtocol) Packet(result result.ResultInter) []byte {
	return nil
}

// 通用解析 ---> 解析为 [][]byte
func commonRespProtocolAnalysis(bs []byte) ([][]byte, error) {
	bsLength := len(bs)
	firstR := findFirstR(bs, 0, bsLength)
	if firstR == -1 {
		return nil, constant.ParamsAnalysisError
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
			return nil, constant.ParamsAnalysisError
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
