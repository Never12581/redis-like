package model

import (
	"fmt"
	"testing"
)

func Test_Sds(t *testing.T) {
	bs := []byte("hello world!")
	s := NewSds(bs)
	s.Append([]byte("\r\nI Loved the world!"))
	fmt.Println(s.Value())
	fmt.Println(string(s.Value()))
}

func Test_No(t *testing.T) {
	bs := []byte("hello world!")
	fmt.Println(string(bs))
	bbs := []byte("\r\nI Loved the world!")
	bbbs := append(bs, bbs...)
	fmt.Println(bbbs)
	s := string(bbbs)
	fmt.Println(s)
}
