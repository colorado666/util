package aslice

import (
	"strconv"
	"strings"
)

//int64 切片 是否包含元素
func ContainInt64(slice []int64, elem int64) bool {
	index := IndexInt64(slice, elem)
	if index >= 0 {
		return true
	}
	return false
}

//int64 切片 元素下标
func IndexInt64(slice []int64, elem int64) int {
	index := -1
	for i, s := range slice {
		if s == elem {
			index = i
			break
		}
	}
	return index
}

//int64 切片 删除元素
func RemoveInt64(slice []int64, elem ...int64) []int64 {
	maps := map[int64]bool{}
	for _, e := range elem {
		maps[e] = true
	}

	var newSlice []int64
	for _, s := range slice {
		if _, ok := maps[s]; !ok {
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

//int64 切片 元素去重复
func DistinctInt64(slice []int64) []int64 {
	maps := map[int64]bool{}
	var newSlice []int64
	for _, s := range slice {
		if _, ok := maps[s]; !ok {
			maps[s] = true
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

//求和
func SumInt64(source []int64) (sum int64) {
	for _, v := range source {
		sum += v
	}
	return
}

//分割字符串为 int64 数组
func SplitToInt64(str string, sep string) (rs []int64, err error) {
	if sep == "" {
		sep = ","
	}
	strs := strings.Split(strings.TrimSpace(str), sep)
	for _, s := range strs {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return rs, err
		}
		rs = append(rs, i)
	}
	return rs, nil
}

//拼接 int64 数组为字符串
func JoinFromInt64(source []int64, sep string) (str string) {
	if sep == "" {
		sep = ","
	}
	for _, i := range source {
		str += strconv.FormatInt(i, 10) + sep
	}
	return strings.TrimSuffix(str, sep)
}
