package protocol

import (
	"fmt"
	"testing"
)

func Test_byte2int(t *testing.T) {
	b := '0'
	i := int(b) - 48
	fmt.Println(b, i)

}
