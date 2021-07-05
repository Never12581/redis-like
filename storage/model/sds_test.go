package model

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestSdsNew(t *testing.T) {
	start := time.Now().Nanosecond()
	sds := SdsNew([]byte("hello world!"))
	fmt.Println(sds.ToString())
	for i := 0; i < 90000; i++ {
		sds.SdsCat([]byte(strconv.Itoa(i)))
		sds.SdsCat([]byte("｜abcdef｜"))
	}
	end := time.Now().Nanosecond()
	fmt.Println("use time : " + strconv.Itoa(end-start))
	//fmt.Println(sds.ToString())

	start = time.Now().Nanosecond()
	bs := []byte("hello world!")
	for i := 0; i < 90000; i++ {
		bs = append(bs, []byte(strconv.Itoa(i))...)
		bs = append(bs, []byte("｜abcdef｜")...)
	}
	end = time.Now().Nanosecond()
	fmt.Println("use time : " + strconv.Itoa(end-start))
	//fmt.Println(string(bs))
	//fmt.Println(len(bs))
	//fmt.Println(cap(bs))
}
