package avalid

import (
	"fmt"
	"gitee.com/asktop_golib/util/astring"
)

//必须为电话号码
type isTel struct {
	isRequired func() bool
	title      string
	value      interface{}
	valueStr   string
	msgs       []string
}

func (c *isTel) Check() (msg string, ok bool) {
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}

	if !c.isRequired() && len(c.valueStr) == 0 {
		return "", true
	}

	if !astring.IsTel(c.valueStr) {
		if len(c.msgs) == 0 {
			msg = fmt.Sprintf("%s 电话号码格式不正确", c.title)
		}
		return msg, false
	}
	return "", true
}
