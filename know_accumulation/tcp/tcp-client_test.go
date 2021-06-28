package tcp

import (
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"testing"
	"time"
)

func TestTcpClient(t *testing.T) {
	for i := 0; i < 1; i++ {
		go call(i)
	}

	time.Sleep(time.Second * 10)
}

func call(i int) {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		//fmt.Println(err.Error())
		panic(err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("maybe,I will fill love with kelly" + strconv.Itoa(i)))
	if err != nil {
		fmt.Println(err.Error() + "25")
	}

	bs := make([]byte, 1024)
	n, err := conn.Read(bs)
	if err != nil {
		//fmt.Println(err.Error() + "31")
		panic(err)
		return
	}

	fmt.Println(string(bs[0:n]))

}

func singleTcpClient() {
	c, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	go func() {
		n, err := c.Write([]byte("from client : I have a dream!\r\n"))
		if err != nil {
			panic(err)
		}
		fmt.Println("n : ", n)
		fmt.Println("=============")
		return
	}()

	rebs, err := ioutil.ReadAll(c)
	if err != nil {
		panic(err)
	}

	fmt.Println("client -> : " + string(rebs))
}
