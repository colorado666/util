package abox

import "gitee.com/asktop_golib/util/acast"

type BoxOptions []BoxOptionItem

type BoxOptionItem struct {
	Name  string      `json:"name"`  //名称（显示用）
	Title string      `json:"title"` //名称（显示用）
	Value interface{} `json:"value"` //数值（唯一）
	BoxItemOther
}

//按code获取选中常量列表（有序）
func (b Box) GetOptionsByCode(checkedCode ...string) BoxOptions {
	boxs := []BoxOptionItem{}
	checkedMap := map[string]bool{}
	for _, v := range checkedCode {
		checkedMap[v] = true
	}
	for _, v := range b {
		box := BoxOptionItem{Name: v.Name, Title: v.Name, Value: v.Code}
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
func (b Box) GetOptionsByNumber(checkedNumber ...int) BoxOptions {
	boxs := []BoxOptionItem{}
	checkedMap := map[int]bool{}
	for _, v := range checkedNumber {
		checkedMap[v] = true
	}
	for _, v := range b {
		box := BoxOptionItem{Name: v.Name, Title: v.Name, Value: v.Number}
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

func (bo BoxOptions) Get(value string) BoxOptionItem {
	for _, v := range bo {
		if acast.ToString(v.Value) == value {
			return v
		}
	}
	return BoxOptionItem{}
}

func (bo BoxOptions) GetMap() map[string]string {
	rs := map[string]string{}
	for _, v := range bo {
		rs[acast.ToString(v.Value)] = v.Name
	}
	return rs
}
