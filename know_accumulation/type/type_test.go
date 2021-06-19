package _type

import (
	"fmt"
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	var i int32
	i = 1
	tt := reflect.TypeOf(i)
	fmt.Println(tt)

	var bs []byte
	bs = []byte("abc")
	ttbs := reflect.TypeOf(bs)
	fmt.Println(ttbs)
}
