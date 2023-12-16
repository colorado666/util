package avalid

import (
	"fmt"
	"gitee.com/asktop_golib/util/acast"
	"github.com/shopspring/decimal"
)

//数值比较
// rs：比较状态 0：等于；1：大于；-1：小于；10：大于等于；-10：小于等于
type cmp struct {
	isRequired func() bool
	title      string
	value      interface{}
	valueStr   string
	msgs       []string
	number     interface{}
	rs         int
}

func (c *cmp) Check() (msg string, ok bool) {
	if len(c.msgs) > 0 {
		msg = c.msgs[0]
	}
	val, err := decimal.NewFromString(c.valueStr)
	if err != nil {
		return fmt.Sprintf("%s非数值", c.title), false
	}

	if !c.isRequired() && val.Equal(decimal.Zero) {
		return "", true
	}

	numberStr := acast.ToString(c.number)
	number, err := decimal.NewFromString(numberStr)
	if err != nil {
		return fmt.Sprintf("%s校验值非数值", c.title), false
	}
	switch c.rs {
	case 0:
		if val.Cmp(number) != 0 {
			if len(c.msgs) == 0 {
				msg = fmt.Sprintf("%s必须等于&s", c.title, numberStr)
			}
			return msg, false
		}
	case 1:
		if val.Cmp(number) != 1 {
			if len(c.msgs) == 0 {
				msg = fmt.Sprintf("%s必须大于&s", c.title, numberStr)
			}
			return msg, false
		}
	case -1:
		if val.Cmp(number) != -1 {
			if len(c.msgs) == 0 {
				msg = fmt.Sprintf("%s必须小于&s", c.title, numberStr)
			}
			return msg, false
		}
	case 10:
		if val.Cmp(number) == -1 {
			if len(c.msgs) == 0 {
				msg = fmt.Sprintf("%s必须大于等于&s", c.title, numberStr)
			}
			return msg, false
		}
	case -10:
		if val.Cmp(number) == 1 {
			if len(c.msgs) == 0 {
				msg = fmt.Sprintf("%s必须小于等于&s", c.title, numberStr)
			}
			return msg, false
		}
	default:
		return "", true
	}
	return "", true
}
