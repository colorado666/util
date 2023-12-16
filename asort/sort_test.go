package asort

import (
	"fmt"
	"sort"
	"testing"
)

func TestRStrings(t *testing.T) {
	a := []string{"b", "a", "d", "c"}
	RStrings(a)
	fmt.Println(a)
}

func TestSortParams(t *testing.T) {
	params := map[string]interface{}{}
	params["c"] = "c"
	params["a"] = "a"
	params["d"] = "d"
	params["b"] = 2.3
	params["e"] = 5.4
	fmt.Println(params)
	fmt.Println(SortParamInterface(params, "&"))
}

type demo struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func TestSortDemo(t *testing.T) {
	var users []demo
	users = append(users, demo{Name: "A", Age: 18})
	users = append(users, demo{Name: "B", Age: 15})
	users = append(users, demo{Name: "C", Age: 20})
	fmt.Println(users)

	//正序
	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	fmt.Println(users)

	//倒序
	sort.Slice(users, func(i, j int) bool {
		return users[i].Age > users[j].Age
	})
	fmt.Println(users)
}
