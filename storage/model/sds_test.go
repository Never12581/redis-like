package model

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSdsNew(t *testing.T) {
	sds := SdsNew([]byte("hello world!"))
	fmt.Println(sds.ToString())
	for i := 0; i < 10000; i++ {
		sds.SdsCat([]byte(strconv.Itoa(i)))
		sds.SdsCat([]byte("｜abcdef｜"))
	}
	fmt.Println(sds.ToString())
	fmt.Println(sds.SdsLen())
	fmt.Println(sds.SdsAvail())
	sds.SdsRange(0, 12)
	fmt.Println(sds.ToString())

	//bs := []byte("hello world!")
	//for i := 0; i < 10000; i++ {
	//	bs = append(bs, []byte(strconv.Itoa(i))...)
	//	bs = append(bs, []byte("｜abcdef｜")...)
	//}
	//fmt.Println(string(bs))
	//fmt.Println(len(bs))
	//fmt.Println(cap(bs))
}
