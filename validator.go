package validator

import (
	"fmt"
	"reflect"
	"strings"
)

/*
1、通过反射拿值
2、通过反射拿类型
3、通过反射拿tag
*/

// Validator 水电费
type Validator struct {
	TagName string
	Sep     string
}

// New Validator
func New(tagName, Sep string) *Validator {
	return &Validator{TagName: tagName, Sep: Sep}
}

// Struct 指针校验结构体
func (v *Validator) Struct(s interface{}) error {
	val := reflect.ValueOf(s)
	elem := reflect.Indirect(val)

	for i := 0; i < elem.NumField(); i++ {
		tagStr := elem.Type().Field(i).Tag.Get(v.TagName)
		if tagStr == "" {
			return nil
		}
		for _, g := range strings.Split(tagStr, v.Sep) {
			var funcName, param string
			ss := strings.Split(g, "=")
			if len(ss) == 2 {
				funcName, param = ss[0], ss[1]
			} else {
				funcName, param = ss[0], ""
			}
			valFunc, g := FuncMap[funcName]
			if !g {
				panic(fmt.Sprintf("tag %s invalid", funcName))
			}
			if !valFunc(elem.Type().Field(i), elem.Field(i), param) {
				return fmt.Errorf("validator %s(%s) invalid", funcName, param)
			}
		}
	}
	return nil
}
