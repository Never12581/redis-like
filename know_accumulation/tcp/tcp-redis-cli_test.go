package tcp

import (
	"fmt"
	"io/ioutil"
	"net"
	"testing"
)

func TestRedisCli(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	s := "*1\r\n$7\r\nCOMMAND\r\n"
	n, err := conn.Write([]byte(s))
	if err != nil {
		panic(err)
	}
	fmt.Printf("写入：%v", n)
	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
}
