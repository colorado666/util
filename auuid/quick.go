package auuid

import "strings"

func GetUUID() string {
	u, _ := NewV4()
	return u.String()
}

func GetUUIDStr() string {
	u, _ := NewV4()
	return strings.ReplaceAll(u.String(), "-", "")
}
