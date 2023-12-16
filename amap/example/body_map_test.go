package example

import (
	"fmt"
	"gitee.com/asktop_golib/util/amap"
	"testing"
)

func TestBodyMap(t *testing.T) {
	bm := amap.NewBodyMap()
	bm.Set("b", "b")
	bm.Set("a", "a")
	bm.Set("2", "2")
	bm.Set("1", "1")
	bm.SetBodyMap("2", func(bm amap.BodyMap) {
		bm.Set("i1", "i1")
		bm.Set("i2", "i2")
		bm.SetBodyMap("i3", func(bm amap.BodyMap) {
			bm.Set("j1", "j1")
		})
	})
	fmt.Println(bm.JsonBody())
	fmt.Println(bm)
	fmt.Println(bm.Delete("2", "i1"))
}
