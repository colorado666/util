package example

import (
	"fmt"
	"gitee.com/asktop_golib/util/areflect"
	"testing"
)

func abc() {

}

func TestGetFuncName(t *testing.T) {
	fmt.Println(areflect.GetFuncAllName(abc))
	fmt.Println(areflect.GetFuncName(abc))
}
