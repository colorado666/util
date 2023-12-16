package example

import (
	"fmt"
	"gitee.com/asktop_golib/util/aslice"
	"testing"
)

func TestContain(t *testing.T) {
	arr := []string{"a", "b", "c"}
	fmt.Println(aslice.ContainString(arr, "a"))
	fmt.Println(aslice.ContainString(arr, "d"))
}

func TestSum(t *testing.T) {
	arr := []string{"0.1", "2.5", "-3.4"}
	fmt.Println(aslice.SumString(arr))
}
