package acast

import "gitee.com/asktop_golib/util/amapstruct"

func MapToStruct(input interface{}, output interface{}) error {
	return amapstruct.Decode(input, output)
}

func StructToMap(input interface{}, output interface{}) error {
	return amapstruct.Decode(input, output)
}
