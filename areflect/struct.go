package areflect

import (
	"reflect"
	"strings"
)

//获取结构体名
func GetStructName(obj interface{}) string {
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}
	objName := objType.Name()
	return objName
}

//获取结构体snake名
func GetStructSnakeName(obj interface{}) string {
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}
	objName := objType.Name()
	return snakeString(objName)
}

/*
获取结构体字段的指定tag名集合，不指定tag或tag为orm时:取字段snake名
@param obj 结构体
@param tag 获取结构体指定标签名
*/
func GetStructFieldNames(obj interface{}, tag ...string) (fieldNames []string) {
	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}
	if objType.Kind() != reflect.Struct {
		return
	}
	var tagName string
	if len(tag) > 0 {
		tagName = tag[0]
	}
	fieldNames = getAnonymousFields(objType, tagName, map[string]string{}, map[string]string{})
	return fieldNames
}

/*
获取结构体字段的指定tag名集合，不指定tag或tag为orm时:取字段snake名
@param obj 结构体
@param tag 获取结构体指定标签名
@param withoutFieldName 不获取的字段名
*/
func GetStructFieldNamesWithoutField(obj interface{}, tag string, withoutFieldName ...string) (fieldNames []string) {
	withoutFieldNameMap := map[string]string{}
	for _, field := range withoutFieldName {
		withoutFieldNameMap[field] = field
	}

	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}
	if objType.Kind() != reflect.Struct {
		return
	}
	fieldNames = getAnonymousFields(objType, tag, withoutFieldNameMap, map[string]string{})
	return fieldNames
}

/*
获取结构体字段的指定tag名集合，不指定tag或tag为orm时:取字段snake名
@param obj 结构体
@param tag 获取结构体指定标签名
@param withoutTagName 不获取的字段tag名
*/
func GetStructFieldNamesWithoutTag(obj interface{}, tag string, withoutTagName ...string) (fieldNames []string) {
	withoutTagNameMap := map[string]string{}
	for _, field := range withoutTagName {
		withoutTagNameMap[field] = field
	}

	objType := reflect.TypeOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}
	if objType.Kind() != reflect.Struct {
		return
	}
	fieldNames = getAnonymousFields(objType, tag, map[string]string{}, withoutTagNameMap)
	return fieldNames
}

// 匿名结构体获取
func getAnonymousFields(objType reflect.Type, tagName string, withoutFieldNameMap map[string]string, withoutTagNameMap map[string]string) (fieldNames []string) {
	for i := 0; i < objType.NumField(); i++ {
		fieldType := objType.Field(i)
		// 匿名结构体获取
		if fieldType.Type.Kind() == reflect.Struct && fieldType.Anonymous {
			fieldNames = append(fieldNames, getAnonymousFields(fieldType.Type, tagName, withoutFieldNameMap, withoutTagNameMap)...)
		} else {
			var fieldName string
			if tagName == "" {
				fieldName = fieldType.Name
				fieldName = snakeString(fieldName)
			} else if tagName == "orm" {
				var ok bool
				fieldName, ok = fieldType.Tag.Lookup(tagName)
				if !ok || fieldName == "-" || fieldName == "_" {
					continue
				}
				fieldName = fieldType.Name
				fieldName = snakeString(fieldName)
			} else {
				fieldName = fieldType.Tag.Get(tagName)
				if fieldName == "-" || fieldName == "_" || fieldName == "" {
					continue
				}
			}
			if _, ok := withoutFieldNameMap[fieldType.Name]; ok {
				continue
			}
			if _, ok := withoutTagNameMap[fieldName]; ok {
				continue
			}
			fieldNames = append(fieldNames, fieldName)
		}
	}
	return fieldNames
}

// snake string, XxYy to xx_yy , XxYY to xx_y_y
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}
