package astring

import "strings"

// 转换成 首字母大写
func ToFirstUpper(s string) string {
	s = strings.TrimSpace(s)
	if s != "" {
		s = strings.ToUpper(s[:1]) + s[1:]
	}
	return s
}

// 转换成 首字母小写
func ToFirstLower(s string) string {
	s = strings.TrimSpace(s)
	if s != "" {
		s = strings.ToLower(s[:1]) + s[1:]
	}
	return s
}

// 转换成 大驼峰命名（UserId）
func ToCamelUpper(s string) string {
	if IsNum_Alpha(s) {
		var rs string
		s = strings.TrimSpace(s)
		es := strings.Split(s, "_")
		for _, e := range es {
			rs += ToFirstUpper(e)
		}
		return rs
	} else {
		return s
	}
}

// 转换成 小驼峰命名（userId）
func ToCamelLower(s string) string {
	return ToFirstLower(ToCamelUpper(s))
}

// 转换成 大下划线命名/蛇形命名（USER_ID）
func ToSnakeUpper(s string) string {
	return strings.ToUpper(ToSnakeLower(s))
}

// 转换成 小下划线命名/蛇形命名（user_id）
func ToSnakeLower(s string) string {
	if IsNum_Alpha(s) {
		var rs string
		l := len(s)
		for i := 0; i < l; i++ {
			e := s[i : i+1]
			if IsUpper(e) {
				e = "_" + strings.ToLower(e)
			}
			rs += e
		}
		rs = strings.TrimPrefix(rs, "_")
		rs = strings.Replace(rs, "__", "_", -1)
		return rs
	} else {
		return s
	}
}

// 将拼接的keys根据kvMap转换成拼接的values
func ToKeysToValuesByKvSep(kvMap map[string]string, keys string, keySep string, valueSep string, force ...bool) string {
	ks := strings.Split(keys, keySep)
	var values []string
	for _, key := range ks {
		value, ok := kvMap[key]
		if ok {
			values = append(values, value)
		} else {
			if len(force) > 0 && force[0] {
				values = append(values, value)
			}
		}
	}
	return strings.Join(values, valueSep)
}

// 将拼接的keys根据kvMap转换成拼接的values
func ToKeysToValuesBySep(kvMap map[string]string, keys string, sep string, force ...bool) string {
	return ToKeysToValuesByKvSep(kvMap, keys, sep, sep, force...)
}

// 将拼接的keys根据kvMap转换成拼接的values
func ToKeysToValues(kvMap map[string]string, keys string, force ...bool) string {
	return ToKeysToValuesBySep(kvMap, keys, ",", force...)
}
