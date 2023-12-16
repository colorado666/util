package abox

import "gitee.com/asktop_golib/util/acast"

type BoxOptionsInt []BoxOptionIntItem

type BoxOptionIntItem struct {
	Name  string `json:"name"`  //名称（显示用）
	Title string `json:"title"` //名称（显示用）
	Value int    `json:"value"` //数值（唯一）
	BoxItemOther
}

//按code获取选中常量列表（有序）
func (b Box) GetOptionsIntByCode(checkedCode ...string) BoxOptionsInt {
	boxs := []BoxOptionIntItem{}
	checkedMap := map[string]bool{}
	for _, v := range checkedCode {
		checkedMap[v] = true
	}
	for _, v := range b {
		box := BoxOptionIntItem{Name: v.Name, Title: v.Name, Value: acast.ToInt(v.Code)}
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
func (b Box) GetOptionsIntByNumber(checkedNumber ...int) BoxOptionsInt {
	boxs := []BoxOptionIntItem{}
	checkedMap := map[int]bool{}
	for _, v := range checkedNumber {
		checkedMap[v] = true
	}
	for _, v := range b {
		box := BoxOptionIntItem{Name: v.Name, Title: v.Name, Value: v.Number}
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

func (bo BoxOptionsInt) Get(value string) BoxOptionIntItem {
	for _, v := range bo {
		if acast.ToString(v.Value) == value {
			return v
		}
	}
	return BoxOptionIntItem{}
}

func (bo BoxOptionsInt) GetMap() map[string]string {
	rs := map[string]string{}
	for _, v := range bo {
		rs[acast.ToString(v.Value)] = v.Name
	}
	return rs
}
