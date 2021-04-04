package main

import (
	"context"
	"flag"
	"github.com/Allenxuxu/gev"
	"github.com/Allenxuxu/gev/connection"
	"github.com/Allenxuxu/ringbuffer"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"redis-like/protocol"
	"strconv"
	"time"
)

type example struct {
	DB *leveldb.DB
}

func (s *example) OnConnect(c *connection.Connection) {
}

func (s *example) OnMessage(c *connection.Connection, ctx interface{}, data []byte) (out []byte) {
	d := time.Now().Add(1000 * time.Millisecond)
	cctx, closeFunc := context.WithDeadline(context.Background(), d)
	defer closeFunc()
	dealCmd, bs := protocol.RespProtocolAnalysis(cctx, data)
	if len(bs) != 0 {
		out = bs
		return
	}
	out = dealCmd.Deal(s.DB)
	return
}

func (s *example) OnClose(c *connection.Connection) {
	log.Println("OnClose ：", c.PeerAddr())
	log.Println("============")
}

func (s *example) UnPacket(c *connection.Connection, buffer *ringbuffer.RingBuffer) (interface{}, []byte) {
	ret := buffer.Bytes()
	buffer.RetrieveAll()
	return nil, ret
}

func (s *example) Packet(c *connection.Connection, data []byte) []byte {
	return append(data, []byte("\r\n")...)
}

func main() {
	DB, err := leveldb.OpenFile("redis_like", nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	handler := &example{DB: DB}

	var port int
	var loops int

	flag.IntVar(&port, "port", 6379, "server port")
	flag.IntVar(&loops, "loops", 0, "num loops")
	flag.Parse()

	s, err := gev.NewServer(handler,
		gev.Network("tcp"),
		gev.Address(":"+strconv.Itoa(port)),
		gev.NumLoops(loops),
		gev.ReusePort(true),
		gev.Protocol(handler),
	)
	if err != nil {
		panic(err)
	}

	s.Start()
}
