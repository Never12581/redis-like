package protocol

import (
	"context"
	"github.com/Allenxuxu/gev/log"
	"redis-like/storage"
	"strconv"
)

var (
	RequestErr      = []byte("-CommonErr request .")
	RequestStartErr = []byte("-CommonErr request not start with * .")
	Pong            = []byte("+Pong")
	OK              = []byte("+OK")
	CommonErr       = []byte("-cannot deal error .")
	UnsupportedErr  = []byte("-the func unsupported .")
	NotFoundErr     = []byte("-the key not found .")

	EmptyCmd     = &Cmd{}
	EmptyCmdDeal = &CmdDeal{}
)

// resp 协议解析
// Cmd：解析出的命令，[]byte：如果解析失败返回的值
func RespProtocolAnalysis(ctx context.Context, db storage.Storage, bs []byte) (*CmdDeal, []byte) {
	cmd, bs := commonRespProtocolAnalysis(bs)
	if len(bs) != 0 {
		return EmptyCmdDeal, bs
	}

	operatorType := string(cmd.ParamBs[0])
	// note : 将初始化方式由 if else 或者 switch case 转化为 map 形式
	// 			降低 时间复杂度
	return NewCmdDeal(GetCmdInitFunc(operatorType)(cmd), db), nil
}

// 通用解析 ---> 解析为 [][]byte
func commonRespProtocolAnalysis(bs []byte) (*Cmd, []byte) {
	bsLength := len(bs)
	firstR := findFirstR(bs, 0, bsLength)
	if firstR == -1 {
		return EmptyCmd, RequestStartErr
	}
	// 本次数组中参数长度
	paramLength, err := strconv.Atoi(string(bs[1:firstR]))
	if err != nil {
		log.Debug(err)
		return EmptyCmd, RequestErr
	}
	// 逐个解析内容
	paramContents := make([][]byte, paramLength)
	tempStart := firstR + 2 // the first position of /n
	tempEnd := bsLength
	for i := 0; i < paramLength; i++ {
		bbs, nextOffset := analysisParamAndNextOffset(bs, tempStart, tempEnd)
		if nextOffset == -1 {
			return EmptyCmd, RequestErr
		}
		tempStart = nextOffset
		paramContents[i] = bbs
	}
	cmd := &Cmd{
		ParamBs: paramContents,
	}
	return cmd, nil
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
