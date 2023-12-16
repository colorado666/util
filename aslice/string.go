package aslice

import "github.com/shopspring/decimal"

//string 切片 是否包含元素
func ContainString(slice []string, elem string) bool {
	index := IndexString(slice, elem)
	if index >= 0 {
		return true
	}
	return false
}

//string 切片 元素下标
func IndexString(slice []string, elem string) int {
	index := -1
	for i, s := range slice {
		if s == elem {
			index = i
			break
		}
	}
	return index
}

//string 切片 删除元素
func RemoveString(slice []string, elem ...string) []string {
	maps := map[string]bool{}
	for _, e := range elem {
		maps[e] = true
	}

	var newSlice []string
	for _, s := range slice {
		if _, ok := maps[s]; !ok {
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

//string 切片 元素去重复
func DistinctString(slice []string) []string {
	maps := map[string]bool{}
	var newSlice []string
	for _, s := range slice {
		if _, ok := maps[s]; !ok {
			maps[s] = true
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

//求和
func SumString(source []string) (sum string, err error) {
	var s decimal.Decimal
	for _, v := range source {
		d, e := decimal.NewFromString(v)
		if e != nil {
			return sum, e
		}
		s = s.Add(d)
	}
	return s.String(), nil
}
