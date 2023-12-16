package avalid

import (
	"fmt"
	"gitee.com/asktop_golib/util/astring"
)

//必须为Email
type isEmail struct {
	isRequired func() bool
	title      string
	value      interface{}
	valueStr   string
	msgs       []string
}

func (c *isEmail) Check() (msg string, ok bool) {
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}

	if !c.isRequired() && len(c.valueStr) == 0 {
		return "", true
	}

	if !astring.IsEmail(c.valueStr) {
		if len(c.msgs) == 0 {
			msg = fmt.Sprintf("%s Email格式不正确", c.title)
		}
		return msg, false
	}
	return "", true
}
