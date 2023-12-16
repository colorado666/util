package example

import (
	"fmt"
	"gitee.com/asktop_golib/util/alimit"
	"testing"
	"time"
)

func TestNewLimit(t *testing.T) {
	limit := alimit.NewLimit(alimit.Level_High)
	apiUniqueKey := "Request.Method" + "Request.URL.Path" + "Input.IP or UserId or Token"
	for i := 1; i <= 100; i++ {
		ok, times := limit.Check(apiUniqueKey, 30, 60)
		if !ok {
			fmt.Println(apiUniqueKey, i, ":", times, "频率受限")
		} else {
			fmt.Println(apiUniqueKey, i, ":", times)
		}
		time.Sleep(time.Second)
	}
}
