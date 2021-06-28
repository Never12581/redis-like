package protocol

import (
	"redis-like/cmd"
	"redis-like/constant"
	"redis-like/executor/result"
	"regexp"
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

func (r *RespProtocol) Packet(rr result.ResultInter) result.ResultInter {
	rbss := make([][]byte, 0)
	if !rr.Success() {
		bs := []byte(rr.Error().Error())
		rbs := make([]byte, 0)
		rbs = append(rbs, '-')
		rbs = append(rbs, bs...)
		rbs = append(rbs, '\r', '\n')
		rbss = append(rbss, []byte{'-'}, bs, []byte{'\r', '\n'})
		r := result.SuccessAndErrorResult(rbss, rr.Error())
		return r
	}
	bss := rr.Result()
	if len(bss) == 1 {
		bs := bss[0]
		// 批量字符串
		if bs == nil {
			rbss = append(rbss, []byte{'$', '-', '1', '\r', '\n'})
			return result.SuccessResult(rbss)
		}
		if len(bs) == 0 {
			rbss = append(rbss, []byte{'$', '0', '\r', '\n', '\r', '\n'})
			return result.SuccessResult(rbss)
		}
		// 数字情况
		if isInt(string(bs)) {
			rbs := make([]byte, 0)
			rbs = append(rbs, ':')
			rbs = append(rbs, bs...)
			rbs = append(rbs, '\r', '\n')
			rbss = append(rbss, rbs)
			return result.SuccessResult(rbss)
		}
		// 均以简单字符串处理
		rbs := make([]byte, 0)
		rbs = append(rbs, '+')
		rbs = append(rbs, bs...)
		rbs = append(rbs, '\r', '\n')
		rbss = append(rbss, rbs)
		return result.SuccessResult(rbss)
	}
	panic("unsupported resp array analysis!")
	return nil
}

// start -------------------------------------------- unPacket 辅助方法
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

// end -------------------------------------------- unPacket 辅助方法

// start -------------------------------------------- packet 辅助方法
// 判断是否为整数
func isInt(s string) bool {
	match, _ := regexp.MatchString(`^[\+-]?\d+$`, s)
	return match
}

// end -------------------------------------------- packet 辅助方法
