package abox

import "gitee.com/asktop_golib/util/acast"

type BoxOptionsString []BoxOptionStringItem

type BoxOptionStringItem struct {
	Name  string `json:"name"`  //名称（显示用）
	Title string `json:"title"` //名称（显示用）
	Value string `json:"value"` //数值（唯一）
	BoxItemOther
}

//按code获取选中常量列表（有序）
func (b Box) GetOptionsStringByCode(checkedCode ...string) BoxOptionsString {
	boxs := []BoxOptionStringItem{}
	checkedMap := map[string]bool{}
	for _, v := range checkedCode {
		checkedMap[v] = true
	}
	for _, v := range b {
		box := BoxOptionStringItem{Name: v.Name, Title: v.Name, Value: v.Code}
		box.BoxItemOther = v.BoxItemOther
		if len(checkedMap) > 0 {
			if _, ok := checkedMap[v.Code]; ok {
				box.Checked = true
			} else {
				box.Checked = false
			}
		}
		boxs = append(boxs, box)
	}
	return boxs
}

//按number获取选中常量列表（有序）
func (b Box) GetOptionsStringByNumber(checkedNumber ...int) BoxOptionsString {
	boxs := []BoxOptionStringItem{}
	checkedMap := map[int]bool{}
	for _, v := range checkedNumber {
		checkedMap[v] = true
	}
	for _, v := range b {
		box := BoxOptionStringItem{Name: v.Name, Title: v.Name, Value: acast.ToString(v.Number)}
		box.BoxItemOther = v.BoxItemOther
		if len(checkedMap) > 0 {
			if _, ok := checkedMap[v.Number]; ok {
				box.Checked = true
			} else {
				box.Checked = false
			}
		}
		boxs = append(boxs, box)
	}
	return boxs
}

func (bo BoxOptionsString) Get(value string) BoxOptionStringItem {
	for _, v := range bo {
		if v.Value == value {
			return v
		}
	}
	return BoxOptionStringItem{}
}

func (bo BoxOptionsString) GetMap() map[string]string {
	rs := map[string]string{}
	for _, v := range bo {
		rs[v.Value] = v.Name
	}
	return rs
}
