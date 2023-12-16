package avalid

import (
	"fmt"
	"gitee.com/asktop_golib/util/aslice"
)

//在切片中
type inSlice struct {
	isRequired func() bool
	title      string
	value      interface{}
	valueStr   string
	msgs       []string
	slice      []string
}

func (c *inSlice) Check() (msg string, ok bool) {
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}

	if !c.isRequired() && len(c.valueStr) == 0 {
		return "", true
	}

	if !aslice.ContainString(c.slice, c.valueStr) {
		if len(c.msgs) == 0 {
			msg = fmt.Sprintf("%s不在规定范围内", c.title)
		}
		return msg, true
	}
	return "", true
}
