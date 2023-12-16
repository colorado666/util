package avalid

import (
	"fmt"
	"gitee.com/asktop_golib/util/astring"
)

//必须为身份证号码
type isIDCard struct {
	isRequired func() bool
	title      string
	value      interface{}
	valueStr   string
	msgs       []string
}

func (c *isIDCard) Check() (msg string, ok bool) {
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}

	if !c.isRequired() && len(c.valueStr) == 0 {
		return "", true
	}

	if !astring.IsIDCard(c.valueStr) {
		if len(c.msgs) == 0 {
			msg = fmt.Sprintf("%s 身份证号码格式不正确", c.title)
		}
		return msg, false
	}
	return "", true
}
