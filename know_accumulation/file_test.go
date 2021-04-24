package know_accumulation

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"testing"
)

func Test_file_fd(t *testing.T) {

	file , err := os.Open("/Users/huangbocai/go/src/t-redis-like/Makefile")
	if err != nil {
		panic(err)
	}

	fmt.Println(file.Fd())

	bs , err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}

func Test_net_fd(t *testing.T) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		fmt.Println(conn)
	}

}
