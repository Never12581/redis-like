package tcp

import (
	"fmt"
	"golang.org/x/sys/unix"
	"io"
	"net"
	"testing"
)

func TestTcpServer(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	listener := l.(*net.TCPListener)
	fl, err := listener.File()
	if err != nil {
		panic(err)
	}

	err = unix.SetNonblock(int(fl.Fd()), true)
	if err != nil {
		panic(err)
	}

	for {
		tc, err := listener.AcceptTCP()
		if err == nil {
			handler(tc)
		} else {
			panic(err)
		}
	}
}

func handler(conn *net.TCPConn) {
	//defer conn.Close()
	buf := make([]byte, 1024)
	//给客户端发送连接成功的信号
	for {
		//持续读取客户端数据，保存在buf缓冲区中，并处理
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("outtoleave", conn.RemoteAddr().String())
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("read err", err)
			return
		}
		//这里是处理数据的一个示范，把客户端发来的数据全部转化为大写
		conn.Write([]byte("+OK\r\n"))
	}
}
