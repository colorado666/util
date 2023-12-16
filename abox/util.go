package abox

import "fmt"

//格式化BoxItem的name
func FormatBoxItemName(boxItem *BoxItem, format string, args ...interface{}) *BoxItem {
	return &BoxItem{Code: boxItem.Code, Number: boxItem.Number, Name: fmt.Sprintf(format, args...)}
}

//格式化BoxItem的name
func FormatBoxItemNameQuick(boxItem *BoxItem, format string) *BoxItem {
	return FormatBoxItemName(boxItem, format, boxItem.Name)
}
