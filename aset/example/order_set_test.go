package example

import (
	"fmt"
	"gitee.com/asktop_golib/util/aset"
	"testing"
)

func TestOrderSet(t *testing.T) {
	set := aset.NewOrderSet()
	set.Add("a")
	set.Add("b")
	set.Add("a")
	set.Set("c")
	set.Set(1)
	fmt.Println(set.Size())
	fmt.Println(set.String())
}

func TestOrderSet2(t *testing.T) {
	arr := []string{
		"a",
		"c",
		"b",
		"a",
	}
	set := aset.NewOrderSetFrom(arr)
	set.Set(1)
	fmt.Println(set.Size())
	fmt.Println(set.String())
	fmt.Println(set.Slice())
	fmt.Println(set.SliceString())
	fmt.Println(set.SliceInt())
}
