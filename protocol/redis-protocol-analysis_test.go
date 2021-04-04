package protocol

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_RespProtocolAnalysis(t *testing.T) {

	bs := []byte("*15\r\n$5\r\nlpush\r\n$1\r\na\r\n$2\r\nc0\r\n$2\r\nc1\r\n$2\r\nc2\r\n$2\r\nc3\r\n$2\r\nc4\r\n$2\r\nc5\r\n$2\r\nc6\r\n$2\r\nc7\r\n$2\r\nc8\r\n$2\r\nc9\r\n$3\r\nc10\r\n$3\r\nc11\r\n$3\r\nc12\r\n")
	cmd, errBs := RespProtocolAnalysis(bs)
	fmt.Println(string(errBs))
	fmt.Println(json.Marshal(cmd))
}
