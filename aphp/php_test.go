package aphp

import (
	"fmt"
	"testing"
)

//php序列化和反序列化
func TestPhpSerialize(t *testing.T) {
	str := `a:1:{s:3:"php";s:24:"世界上最好的语言";}`

	// unserialize() in php
	out := map[interface{}]interface{}{}
	err := Unmarshal([]byte(str), &out)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(out) //map[php:世界上最好的语言]

	// serialize() in php
	jsonbyte, _ := Marshal(out, nil)
	t.Log(string(jsonbyte)) // a:1:{s:3:"php";s:24:"世界上最好的语言";}
}

func TestPhpSerialize2(t *testing.T) {
	data := new(TestPhp)
	data.PhpLang = "世界上最好的语言"

	// serialize() in php
	jsonbyte, _ := Marshal(data, nil)
	t.Log(string(jsonbyte)) // O:7:"TestPhp":1:{s:8:"php_lang";s:24:"世界上最好的语言";}

	str := `O:7:"TestPhp":1:{s:8:"php_lang";s:24:"世界上最好的语言";}`

	out := new(TestPhp)
	err := Unmarshal([]byte(str), out)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(out) //map[php:世界上最好的语言]
}

type TestPhp struct {
	PhpLang string `php:"php_lang"`
}
