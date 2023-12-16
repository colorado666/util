package arsa

import (
    "fmt"
    "testing"
    "time"
)

func TestGenRsaKey(t *testing.T) {
    //private, public, err := GenRsaKey(1024)
    private, public, err := GenRsaKey(2048)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(private)
    fmt.Println(public)

    time.Sleep(time.Second)
}
