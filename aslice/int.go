package aslice

import (
	"strconv"
	"strings"
)

//int 切片 是否包含元素
func ContainInt(slice []int, elem int) bool {
	index := IndexInt(slice, elem)
	if index >= 0 {
		return true
	}
	return false
}

//int 切片 元素下标
func IndexInt(slice []int, elem int) int {
	index := -1
	for i, s := range slice {
		if s == elem {
			index = i
			break
		}
	}
	return index
}

//int 切片 删除元素
func RemoveInt(slice []int, elem ...int) []int {
	maps := map[int]bool{}
	for _, e := range elem {
		maps[e] = true
	}

	var newSlice []int
	for _, s := range slice {
		if _, ok := maps[s]; !ok {
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

//int 切片 元素去重复
func DistinctInt(slice []int) []int {
	maps := map[int]bool{}
	var newSlice []int
	for _, s := range slice {
		if _, ok := maps[s]; !ok {
			maps[s] = true
			newSlice = append(newSlice, s)
		}
	}
	return newSlice
}

//求和
func SumInt(source []int) (sum int) {
	for _, v := range source {
		sum += v
	}
	return
}

//分割字符串为 int 数组
func SplitToInt(str string, sep string) (rs []int, err error) {
	if sep == "" {
		sep = ","
	}
	strs := strings.Split(strings.TrimSpace(str), sep)
	for _, s := range strs {
		i, err := strconv.Atoi(s)
		if err != nil {
			return rs, err
		}
		rs = append(rs, i)
	}
	return rs, nil
}

//拼接 int 数组为字符串
func JoinFromInt(source []int, sep string) (str string) {
	if sep == "" {
		sep = ","
	}
	for _, i := range source {
		str += strconv.Itoa(i) + sep
	}
	return strings.TrimSuffix(str, sep)
}
