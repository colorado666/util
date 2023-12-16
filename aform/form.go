package aform

import (
	"encoding/json"
	"gitee.com/asktop_golib/util/abox"
	"gitee.com/asktop_golib/util/acast"
	"github.com/shopspring/decimal"
)

type FormItemType string

const (
	//配置值类型
	FormItemType_text     FormItemType = "text"     //字符串
	FormItemType_textarea FormItemType = "textarea" //多行文本
	FormItemType_richtext FormItemType = "richtext" //富文本
	FormItemType_int      FormItemType = "int"      //整型（go用int64接收，若前端显示需要小数处理，需设置decimal）
	FormItemType_number   FormItemType = "number"   //数值（go用string接收，可以是浮点数值）
	FormItemType_radio    FormItemType = "radio"    //单选框
	FormItemType_checkbox FormItemType = "checkbox" //复选框
	FormItemType_image    FormItemType = "image"    //图片链接
	FormItemType_file     FormItemType = "file"     //文件链接
)

//通用表单
type Form []FormItem

//配置参数
type FormItem struct {
	Title    string               `json:"title,omitempty"`     //配置名
	Name     string               `json:"name"`                //配置标识（唯一）
	Value    interface{}          `json:"value"`               //配置值
	Type     FormItemType         `json:"type,omitempty"`      //配置值类型
	Tips     string               `json:"tips,omitempty"`      //配置说明
	Decimal  int                  `json:"decimal,omitempty"`   //前端显示处理小数位数（只有int时设置）
	Options  []abox.BoxOptionItem `json:"options,omitempty"`   //选择项目
	Required bool                 `json:"required,omitempty"`  //是否必须
	Readonly bool                 `json:"readonly,omitempty"`  //是否只读
	Hide     bool                 `json:"hide,omitempty"`      //是否隐藏（不在页面显示）
	ParamKey bool                 `json:"param_key,omitempty"` //是否参数唯一标识
}

func NewForm(formItems ...FormItem) Form {
	if len(formItems) == 0 {
		return []FormItem{}
	}
	return formItems
}

func NewFormFromJson(formItemsJson string) (Form, error) {
	formItems := NewForm()
	err := json.Unmarshal([]byte(formItemsJson), &formItems)
	return formItems, err
}

func NewFormFromMap(paramMap map[string]interface{}) Form {
	formItems := NewForm()
	for key, val := range paramMap {
		formItem := FormItem{
			Name:  key,
			Value: val,
		}
		formItems = append(formItems, formItem)
	}
	return formItems
}

func NewFormFromMapJson(paramMapJson string) (Form, error) {
	paramMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(paramMapJson), &paramMap)
	if err != nil {
		return nil, err
	}
	formItems := NewFormFromMap(paramMap)
	return formItems, nil
}

//对象map转换为配置
func (m Form) FromMap(paramMap map[string]interface{}) Form {
	if paramMap == nil {
		return m
	}
	paramForm := NewFormFromMap(paramMap)
	paramForm = paramForm.FilterToDefault(m)
	return paramForm
}

//对象json转换为配置
func (m Form) FromMapJson(paramMapJson string) (Form, error) {
	if paramMapJson == "" {
		return m, nil
	}
	paramForm, err := NewFormFromMapJson(paramMapJson)
	if err != nil {
		return nil, err
	}
	paramForm = paramForm.FilterToDefault(m)
	return paramForm, nil
}

//转换为Map
func (m Form) ToMap() map[string]interface{} {
	paramMap := map[string]interface{}{}
	for _, item := range m {
		paramMap[item.Name] = item.Value
	}
	return paramMap
}

//转换为对象
func (m Form) ToObject(obj interface{}) error {
	paramMap := m.ToMap()
	return acast.MapToStruct(paramMap, obj)
}

//过滤处理配置（删除无效配置，添加未设置配置为默认配置）
func (m Form) FilterToDefault(defaultForm Form) Form {
	formItemMap := map[string]FormItem{}
	for _, item := range m {
		formItemMap[item.Name] = item
	}

	paramForm := NewForm()
	for _, defaultItem := range defaultForm {
		if item, ok := formItemMap[defaultItem.Name]; ok {
			if !defaultItem.Readonly {
				defaultItem.Value = item.Value
			}
		}
		paramForm = append(paramForm, defaultItem)
	}
	return paramForm
}

func (m Form) Add(item FormItem) Form {
	m = append(m, item)
	return m
}

func (m Form) Set(item FormItem) Form {
	for i, v := range m {
		if v.Name == item.Name {
			m[i] = item
			return m
		}
	}
	return m.Add(item)
}

func (m Form) Get(name string) (FormItem, bool) {
	for _, v := range m {
		if v.Name == name {
			return v, true
		}
	}
	return FormItem{}, false
}

func (m Form) GetString(name string) (string, bool) {
	for _, v := range m {
		if v.Name == name {
			return acast.ToString(v.Value), true
		}
	}
	return "", false
}

func (m Form) GetDecimal(name string) (decimal.Decimal, bool) {
	v, ok := m.GetString(name)
	value, _ := decimal.NewFromString(v)
	return value, ok
}

func (m Form) Del(name string) Form {
	nm := NewForm()
	for _, v := range m {
		if v.Name != name {
			nm.Add(v)
		}
	}
	return nm
}
