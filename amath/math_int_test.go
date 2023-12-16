package amath

import (
    "fmt"
    "testing"
)

func TestInt(t *testing.T) {
    fmt.Println(IntAbs(-12))
    fmt.Println(IntToUint(12))
    fmt.Println(IntToUint(0))
    fmt.Println(IntToUint(-12))
}
