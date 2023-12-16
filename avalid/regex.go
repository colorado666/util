package avalid

import (
	"fmt"
	"gitee.com/asktop_golib/util/astring"
)

//正则表达式验证
type regex struct {
	isRequired func() bool
	title      string
	value      interface{}
	valueStr   string
	msgs       []string
	exp        string
}

func (c *regex) Check() (msg string, ok bool) {
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}

	if !c.isRequired() && len(c.valueStr) == 0 {
		return "", true
	}

	if !astring.MatchString(c.exp, c.valueStr) {
		if len(c.msgs) > 0 {
			msg = fmt.Sprintf("%s验证不合法", c.title)
		}
		return msg, false
	}
	return "", true
}
