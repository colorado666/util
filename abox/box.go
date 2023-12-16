package abox

import (
	"encoding/json"
	"gitee.com/asktop_golib/util/acast"
	"sort"
)

//公用配置项
type Box []*BoxItem

//单项配置
type BoxItem struct {
	Name         string `json:"name"`   //名称（显示用）
	Number       int    `json:"number"` //数值（唯一）
	Code         string `json:"code"`   //标识（唯一）
	Title        string `json:"title"`  //名称（显示用）
	BoxItemOther                        //单项配置扩展
}

//单项配置扩展
type BoxItemOther struct {
	Sort    int    `json:"sort"`    //排序（越小越靠前）
	Color   string `json:"color"`   //颜色
	Checked bool   `json:"checked"` //是否选中
}

//创建
func NewBox(item ...*BoxItem) Box {
	box := []*BoxItem{}
	if len(item) > 0 {
		box = append(box, item...)
	}
	return box
}

//从字符串创建
func NewBoxFromStr(itemsStr string) (Box, error) {
	box := NewBox()
	err := json.Unmarshal([]byte(itemsStr), &box)
	return box, err
}

//从Map创建
func NewBoxFromMap(itemsMap map[string]string) Box {
	box := NewBox()
	for k, v := range itemsMap {
		item := new(BoxItem)
		item.Name = v
		item.Code = k
		item.Number = acast.ToInt(k)
	}
	return box
}

//复制
func (b Box) Copy() Box {
	box := NewBox()
	for _, item := range b {
		box = append(box, item.Copy())
	}
	return box
}

//新增配置项（按数值去重）
func (b Box) AddByNumber(item ...*BoxItem) Box {
	box := NewBox()
	items := map[int]*BoxItem{}
	for _, v := range b {
		box = append(box, v)
		items[v.Number] = v
	}
	for _, v := range item {
		if _, ok := items[v.Number]; !ok {
			box = append(box, v)
			items[v.Number] = v
		}
	}
	return box
}

//新增配置项（按标识去重）
func (b Box) AddByCode(item ...*BoxItem) Box {
	box := NewBox()
	items := map[string]*BoxItem{}
	for _, v := range b {
		box = append(box, v)
		items[v.Code] = v
	}
	for _, v := range item {
		if _, ok := items[v.Code]; !ok {
			box = append(box, v)
			items[v.Code] = v
		}
	}
	return box
}

//在前面新增配置项（按数值去重）
func (b Box) AddFrontByNumber(item ...*BoxItem) Box {
	box := NewBox()
	items := map[int]*BoxItem{}
	for _, v := range item {
		box = append(box, v)
		items[v.Number] = v
	}
	for _, v := range b {
		if _, ok := items[v.Number]; !ok {
			box = append(box, v)
			items[v.Number] = v
		}
	}
	return box
}

//在前面新增配置项（按标识去重）
func (b Box) AddFrontByCode(item ...*BoxItem) Box {
	box := NewBox()
	items := map[string]*BoxItem{}
	for _, v := range item {
		box = append(box, v)
		items[v.Code] = v
	}
	for _, v := range b {
		if _, ok := items[v.Code]; !ok {
			box = append(box, v)
			items[v.Code] = v
		}
	}
	return box
}

//按number删除
func (b Box) DelByNumber(item ...*BoxItem) Box {
	items := map[int]*BoxItem{}
	for _, v := range item {
		items[v.Number] = v
	}
	box := NewBox()
	for _, v := range b {
		if _, ok := items[v.Number]; !ok {
			box = append(box, v)
		}
	}
	return box
}

//按code删除
func (b Box) DelByCode(item ...*BoxItem) Box {
	items := map[string]*BoxItem{}
	for _, v := range item {
		items[v.Code] = v
	}
	box := NewBox()
	for _, v := range b {
		if _, ok := items[v.Code]; !ok {
			box = append(box, v)
		}
	}
	return box
}

//按name获取
func (b Box) GetByName(name string) *BoxItem {
	for _, v := range b {
		if v.Name == name {
			return v
		}
	}
	return nil
}

//按number获取
func (b Box) GetByNumber(number int) *BoxItem {
	for _, v := range b {
		if v.Number == number {
			return v
		}
	}
	return nil
}

//按code获取
func (b Box) GetByCode(code string) *BoxItem {
	for _, v := range b {
		if v.Code == code {
			return v
		}
	}
	return nil
}

//获取codes
func (b Box) GetNames() []string {
	vals := []string{}
	for _, item := range b {
		vals = append(vals, item.Name)
	}
	return vals
}

//获取codes
func (b Box) GetCodes() []string {
	vals := []string{}
	for _, item := range b {
		vals = append(vals, item.Code)
	}
	return vals
}

//获取numbers
func (b Box) GetNumbers() []int {
	vals := []int{}
	for _, item := range b {
		vals = append(vals, item.Number)
	}
	return vals
}

//获取json字符串
func (b Box) GetJsonStr() (string, error) {
	settingByte, err := json.Marshal(b)
	return string(settingByte), err
}

//获取Map
func (b Box) GetMapNumber() map[int]string {
	rs := map[int]string{}
	for _, item := range b {
		rs[item.Number] = item.Name
	}
	return rs
}

//获取Map
func (b Box) GetMapNumberStr() map[string]string {
	rs := map[string]string{}
	for _, item := range b {
		rs[acast.ToString(item.Number)] = item.Name
	}
	return rs
}

//获取Map
func (b Box) GetMapCode() map[string]string {
	rs := map[string]string{}
	for _, item := range b {
		rs[item.Code] = item.Name
	}
	return rs
}

//获取Map
func (b Box) GetMapCodeInt() map[int]string {
	rs := map[int]string{}
	for _, item := range b {
		rs[acast.ToInt(item.Code)] = item.Name
	}
	return rs
}

//获取Map
func (b Box) GetMapNameNumber() map[string]int {
	rs := map[string]int{}
	for _, item := range b {
		rs[item.Name] = item.Number
	}
	return rs
}

//获取Map
func (b Box) GetMapNameNumberStr() map[string]string {
	rs := map[string]string{}
	for _, item := range b {
		rs[item.Name] = acast.ToString(item.Number)
	}
	return rs
}

//获取Map
func (b Box) GetMapNameCode() map[string]string {
	rs := map[string]string{}
	for _, item := range b {
		rs[item.Name] = item.Code
	}
	return rs
}

//获取Map
func (b Box) GetMapNameCodeInt() map[string]int {
	rs := map[string]int{}
	for _, item := range b {
		rs[item.Name] = acast.ToInt(item.Code)
	}
	return rs
}

//处理
func (b Box) FuncBox(funcBox func(box []*BoxItem) Box) Box {
	return funcBox(b)
}

//处理
func (b Box) FuncBoxItem(funcBoxItem func(item *BoxItem) *BoxItem) Box {
	box := NewBox()
	for _, item := range b {
		box = append(box, funcBoxItem(item))
	}
	return box
}

func (s Box) Len() int           { return len(s) }
func (s Box) Less(i, j int) bool { return s[i].Sort < s[j].Sort }
func (s Box) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

//正序排序
func (b Box) Sort() Box {
	sort.Sort(b)
	return b
}

//倒序排序
func (b Box) RSort() Box {
	sort.Slice(b, func(i, j int) bool {
		return b[i].Sort > b[j].Sort
	})
	return b
}
