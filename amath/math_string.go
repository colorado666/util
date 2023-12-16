package amath

import (
    "gitee.com/asktop_golib/util/acast"
    "github.com/shopspring/decimal"
    "strings"
)

// 天花板数，数值变大
func StrCeil(numStr interface{}, n int) string {
    num, err := decimal.NewFromString(acast.ToString(numStr))
    if err != nil {
        return "0." + strings.Repeat("0", n)
    }
    return DecimalCeil(num, n).StringFixed(int32(n))
}

// 地板数，数值变小
func StrFloor(numStr interface{}, n int) string {
    num, err := decimal.NewFromString(acast.ToString(numStr))
    if err != nil {
        return "0." + strings.Repeat("0", n)
    }
    return DecimalFloor(num, n).StringFixed(int32(n))
}

// 四舍五入，正数时数值舍变小入变大，负数时数值舍变大入变小
func StrRound(numStr interface{}, n int) string {
    num, err := decimal.NewFromString(acast.ToString(numStr))
    if err != nil {
        return "0." + strings.Repeat("0", n)
    }
    return DecimalRound(num, n).StringFixed(int32(n))
}

// 直接舍去，正数时数值变小，负数时数值变大
func StrScale(numStr interface{}, n int) string {
    num, err := decimal.NewFromString(acast.ToString(numStr))
    if err != nil {
        return "0." + strings.Repeat("0", n)
    }
    return DecimalScale(num, n).StringFixed(int32(n))
}

// 取绝对值
func StrAbs(numStr string) (string, error) {
    num, err := decimal.NewFromString(numStr)
    if err != nil {
        return "", err
    }
    return num.Abs().String(), nil
}

// 取整
func StrIntPart(numStr string) int64 {
    num, err := decimal.NewFromString(numStr)
    if err != nil {
        return 0
    }
    return num.IntPart()
}

// 次方
func StrShift(numStr string, shift int64) decimal.Decimal {
    num, _ := decimal.NewFromString(numStr)
    return num.Shift(int32(shift))
}

// 0值判读
func StrIsZero(numStr string) bool {
    num, err := decimal.NewFromString(numStr)
    if err != nil {
        return false
    }
    return num.IsZero()
}

func StrAdd(v1 string, v2 string) (string, error) {
    v := "0"

    tmpD1, err := decimal.NewFromString(v1)
    if err != nil {
        return "", err
    }

    tmpD2, err := decimal.NewFromString(v2)
    if err != nil {
        return "", err
    }

    tmpResult := tmpD1.Add(tmpD2)
    v = tmpResult.String()
    return v, nil
}

func StrSub(v1 string, v2 string) (string, error) {
    v := "0"

    tmpD1, err := decimal.NewFromString(v1)
    if err != nil {
        return "", err
    }

    tmpD2, err := decimal.NewFromString(v2)
    if err != nil {
        return "", err
    }

    tmpResult := tmpD1.Sub(tmpD2)
    v = tmpResult.String()
    return v, nil
}

func StrMul(v1 string, v2 string) (string, error) {
    v := "0"

    tmpD1, err := decimal.NewFromString(v1)
    if err != nil {
        return "", err
    }

    tmpD2, err := decimal.NewFromString(v2)
    if err != nil {
        return "", err
    }

    tmpResult := tmpD1.Mul(tmpD2)
    v = tmpResult.String()
    return v, nil
}

func StrDiv(v1 string, v2 string) (string, error) {
    v := "0"

    tmpD1, err := decimal.NewFromString(v1)
    if err != nil {
        return "", err
    }

    tmpD2, err := decimal.NewFromString(v2)
    if err != nil {
        return "", err
    }

    tmpResult := tmpD1.Div(tmpD2)
    v = tmpResult.String()
    return v, nil
}
