package akey

import (
	"fmt"
	"testing"
)

func TestGetKey(t *testing.T) {
	source := "abc"
	fmt.Println(len([]byte(source)))
	key := GetKey(source, 8)
	fmt.Println(len(key))
	fmt.Println(key)
	fmt.Println(string(key))
}
