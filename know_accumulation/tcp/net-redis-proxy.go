package tcp

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

var (
	// 这种地方需要连接池
	proxyConns []net.Conn
)

func init() {
	for i := 0; i < 5; i++ {
		pc, err := net.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			panic(err)
		}
		proxyConns = append(proxyConns, pc)
	}
}

func getProxyConn() net.Conn {
	return proxyConns[rand.Intn(5)]
}

func process(conn net.Conn) {
	// 处理完关闭连接
	defer conn.Close()

	// 针对当前连接做发送和接受操作
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			//fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}

		s := string(buf[:n])
		if strings.Contains(strings.ToLower(s), "command") {
			conn.Write([]byte("+OK\r\n"))
			continue
		}
		fmt.Printf("收到的数据：%v\n", string(buf[:n]))

		pconn := getProxyConn()
		_, err = pconn.Write(buf[:n])
		if err != nil {
			fmt.Printf("write to redis conn failed, err:%v\n", err)
			break
		}

		//var bbf [128]byte
		n, err = pconn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from proxy redis conn failed, err:%v\n", err)
			break
		}

		fmt.Printf("返回的数据：%v\n", string(buf[:n]))
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Printf("write from conn failed, err:%v\n", err)
			break
		}
	}
}

func ProxyStart() {
	// 建立 tcp 服务
	listen, err := net.Listen("tcp", "127.0.0.1:6380")
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}

	go func() {
		for {
			time.Sleep(100 * time.Second)
			for _, pc := range proxyConns {
				n, err := pc.Write([]byte("*1\r\n$4\r\nping\r\n"))
				if err != nil {
					log.Println(pc.RemoteAddr().String() + " write:" + err.Error())
				} else {
					log.Printf("发送字节数：%v", n)
				}
				bss := make([]byte, 128)
				n, err = pc.Read(bss)
				if err != nil {
					log.Println(pc.RemoteAddr().String() + " read:" + err.Error())
				} else {
					log.Printf("返回值：%v", string(bss[:n]))
				}

			}
		}
	}()

	for {
		// 等待客户端建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		// 启动一个单独的 goroutine 去处理连接
		go process(conn)
	}
}
