package avalid

import (
	"fmt"
	"gitee.com/asktop_golib/util/astring"
)

//必须为整数
type isInt struct {
	isRequired func() bool
	title      string
	value      interface{}
	valueStr   string
	msgs       []string
}

func (c *isInt) Check() (msg string, ok bool) {
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}

	if !c.isRequired() && len(c.valueStr) == 0 {
		return "", true
	}

	if !astring.IsInt(c.valueStr) {
		if len(c.msgs) == 0 {
			msg = fmt.Sprintf("%s必须为整数", c.title)
		}
		return msg, false
	}
	return "", true
}
