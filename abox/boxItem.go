package abox

//复制
func (m *BoxItem) Copy() *BoxItem {
	m1 := *m
	m2 := &m1
	return m2
}

//配置项是否相同（同时判断Number和Code）
func (m *BoxItem) Equal(item *BoxItem) bool {
	if item == nil {
		return false
	}
	return m.Number == item.Number && m.Code == item.Code
}

//配置项是否相同（只判断Name）
func (m *BoxItem) EqualName(item *BoxItem) bool {
	if item == nil {
		return false
	}
	return m.Name == item.Name
}

//配置项是否相同（只判断Number）
func (m *BoxItem) EqualNumber(item *BoxItem) bool {
	if item == nil {
		return false
	}
	return m.Number == item.Number
}

//配置项是否相同（只判断Code）
func (m *BoxItem) EqualCode(item *BoxItem) bool {
	if item == nil {
		return false
	}
	return m.Code == item.Code
}

//配置项是否相同（只判断name）
func (m *BoxItem) EqualByName(name string) bool {
	return m.Name == name
}

//配置项是否相同（只判断Number）
func (m *BoxItem) EqualByNumber(number int) bool {
	return m.Number == number
}

//配置项是否相同（只判断Code）
func (m *BoxItem) EqualByCode(code string) bool {
	return m.Code == code
}
