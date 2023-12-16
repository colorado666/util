package amath

import (
    "github.com/shopspring/decimal"
    "strconv"
)

// 直接舍去
func FloatScaleToStringFixed(value float64, precision int) string {
    return strconv.FormatFloat(value, 'f', precision, 64)
}

func FloatScale(value float64, precision int) float64 {
    value, _ = strconv.ParseFloat(FloatScaleToStringFixed(value, precision), 64)
    return value
}

func FloatScaleStringFixed(value string, precision int) string {
    f, _ := strconv.ParseFloat(value, 64)
    return FloatScaleToStringFixed(f, precision)
}

func NewDecimalFromFloatScale(value float64, precision int) decimal.Decimal {
    d, _ := decimal.NewFromString(FloatScaleToStringFixed(value, precision))
    return d
}

// 四舍五入
func FloatRoundToStringFixed(value float64, precision int) string {
    return decimal.NewFromFloat(value).StringFixed(int32(precision))
}

func FloatRound(value float64, precision int) float64 {
    value, _ = strconv.ParseFloat(FloatRoundToStringFixed(value, precision), 64)
    return value
}

func FloatRoundStringFixed(value string, precision int) string {
    f, _ := strconv.ParseFloat(value, 64)
    return FloatRoundToStringFixed(f, precision)
}

func NewDecimalFromFloatRound(value float64, precision int) decimal.Decimal {
    d, _ := decimal.NewFromString(FloatRoundToStringFixed(value, precision))
    return d
}
