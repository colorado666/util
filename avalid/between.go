package avalid

import (
	"fmt"
	"gitee.com/asktop_golib/util/acast"
	"github.com/shopspring/decimal"
)

//数值的范围
type between struct {
	isRequired func() bool
	title      string
	value      interface{}
	valueStr   string
	msgs       []string
	min        interface{}
	max        interface{}
}

func (c *between) Check() (msg string, ok bool) {
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

	if c.min != nil {
		mi, _ := decimal.NewFromString(acast.ToString(c.min))
		if val.Cmp(mi) < 0 {
			if len(c.msgs) == 0 {
				msg = fmt.Sprintf("%s必须大于等于 %s", c.title, mi.String())
			}
			return msg, false
		}
	}
	if c.max != nil {
		ma, _ := decimal.NewFromString(acast.ToString(c.max))
		if val.Cmp(ma) > 0 {
			if len(c.msgs) == 0 {
				msg = fmt.Sprintf("%s必须小于等于 %s", c.title, ma.String())
			}
			return msg, false
		}
	}
	return "", true
}
