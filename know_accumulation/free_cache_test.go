package know_accumulation

import (
	"fmt"
	"github.com/coocood/freecache"
	"testing"
	"time"
)

var cache *freecache.Cache

func init() {
	cacheSize := 100 * 1024 * 1024
	cache = freecache.NewCache(cacheSize)
}

func TestSet(t *testing.T) {
	err := cache.Set([]byte("a"), []byte("b"), 5)
	if err != nil {
		panic(err)
	}
	val, err := cache.Get([]byte("a"))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(val))
	time.Sleep(time.Second * 5)
	val, err = cache.Get([]byte("a"))
	if err != nil {
		panic(err)
	}
	fmt.Println("====" + string(val))
}

func TestDel(t *testing.T) {
	err := cache.Set([]byte("a"), []byte("b"), 5)
	if err != nil {
		panic(err)
	}
	val, err := cache.Get([]byte("a"))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(val))

	cache.Del([]byte("a"))

	val, err = cache.Get([]byte("a"))
	if err != nil {
		panic(err)
	}
	fmt.Println("====" + string(val))
}
