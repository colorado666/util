package example

import (
	"fmt"
	"gitee.com/asktop_golib/util/abox"
	"gitee.com/asktop_golib/util/ajson"
	"testing"
)

func TestBox(t *testing.T) {
	boxA := &abox.BoxItem{Name: "NameA", Code: "A", Number: 1}
	boxB := &abox.BoxItem{Name: "NameB", Code: "B", Number: 2}
	box := abox.NewBox(boxA, boxB)
	options := box.GetOptionsByCode("B")
	fmt.Println(ajson.Encode(options))
	fmt.Println(ajson.Encode(options.GetMap()))
}

func TestBox2(t *testing.T) {
	boxA := &abox.BoxItem{Name: "NameA", Code: "A", Number: 1, BoxItemOther: abox.BoxItemOther{Sort: 2}}
	boxB := &abox.BoxItem{Name: "NameB", Code: "B", Number: 2, BoxItemOther: abox.BoxItemOther{Sort: 1}}
	box := abox.NewBox(boxA, boxB)

	options := box.GetOptionsByCode("B")
	fmt.Println(ajson.Encode(options))

	box.Sort()
	options = box.GetOptionsByCode("B")
	fmt.Println(ajson.Encode(options))

	options = box.Copy().RSort().GetOptionsByCode("B")
	fmt.Println(ajson.Encode(options))

	options = box.GetOptionsByCode("B")
	fmt.Println(ajson.Encode(options))
}
