package model

import (
	"fmt"
	"strconv"
	"testing"
)

func TestList_SearchIndex(t *testing.T) {
	list := List{}
	for i := 0; i < 5; i++ {
		list.AddNodeHead(strconv.Itoa(i))
	}
	list.AddNodeIndex(20, 1, false)
	fmt.Println(list.SearchIndex(0).value)
	fmt.Println(list.SearchIndex(1).value)
	fmt.Println(list.SearchIndex(2).value)
	fmt.Println(list.SearchIndex(3).value)
	fmt.Println(list.SearchIndex(4).value)
	fmt.Println(list.SearchIndex(5).value)
}
