package akey

import "strings"

func GetKey(keySource string, lenght int) (key []byte) {
	suffix := "0000000000000000000000000000000000000000000000000000000000000000"
	keySource = strings.TrimSpace(keySource) + suffix
	key = []byte(keySource)[:lenght]
	return
}
