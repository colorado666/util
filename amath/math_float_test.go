package amath

import (
    "fmt"
    "github.com/shopspring/decimal"
    "testing"
)

func TestFloatScale(t *testing.T) {
    var a float64 = -1.2345
    //fmt.Println(a)
    fmt.Println(FloatScale(a, 3))
    fmt.Println(FloatScale(a, 5))
    fmt.Println(FloatScaleToStringFixed(a, 3))
    fmt.Println(FloatScaleToStringFixed(a, 5))
    fmt.Println(FloatScaleStringFixed("-1.2345", 3))
    fmt.Println(FloatScaleStringFixed("-1.2345", 5))
    fmt.Println(NewDecimalFromFloatScale(a, 5))
    fmt.Println(NewDecimalFromFloatScale(a, 5).Float64())
    fmt.Println(NewDecimalFromFloatScale(a, 5).String())
    fmt.Println(decimal.NewFromFloat(a).StringFixed(3))
    fmt.Println(decimal.NewFromFloat(a).StringFixed(5))
}

func TestFloatRound(t *testing.T) {
    var a float64 = -1.2345
    //fmt.Println(a)
    fmt.Println(FloatRound(a, 3))
    fmt.Println(FloatRound(a, 5))
    fmt.Println(FloatRoundToStringFixed(a, 3))
    fmt.Println(FloatRoundToStringFixed(a, 5))
    fmt.Println(FloatRoundStringFixed("-1.2345", 3))
    fmt.Println(FloatRoundStringFixed("-1.2345", 5))
    fmt.Println(NewDecimalFromFloatRound(a, 5))
    fmt.Println(NewDecimalFromFloatRound(a, 5).String())
    fmt.Println(NewDecimalFromFloatRound(a, 5).Float64())
    fmt.Println(decimal.NewFromFloat(a).StringFixed(3))
    fmt.Println(decimal.NewFromFloat(a).StringFixed(5))
}

