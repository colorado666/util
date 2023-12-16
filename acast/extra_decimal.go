package acast

import (
    "github.com/shopspring/decimal"
)

func ToDecimal(i interface{}, defaultVal ...string) decimal.Decimal {
    v, err := decimal.NewFromString(ToString(i))
    if len(defaultVal) > 0 && err != nil {
        v, _ = decimal.NewFromString(defaultVal[0])
    }
    return v
}
