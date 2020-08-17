package validator

import (
	"reflect"
	"strconv"
)

// FuncMap 校验函数定义map
var FuncMap = map[string]func(reflect.StructField, reflect.Value, string) bool{
	"min": IsMin,
	"max": IsMax,
	"gte": IsGte,
	"lte": isLte,
}

// IsMin min定义
func IsMin(f reflect.StructField, v reflect.Value, param string) bool {
	if f.Type.Kind() != reflect.String {
		panic("IsMin only supports string type")
	}
	l, err := strconv.Atoi(param)
	if err != nil {
		panic("min=.(int)")
	}
	return len(v.Interface().(string)) >= l
}

// IsMax min定义
func IsMax(f reflect.StructField, v reflect.Value, param string) bool {
	if f.Type.Kind() != reflect.String {
		panic("IsMax only supports string type")
	}
	l, err := strconv.Atoi(param)
	if err != nil {
		panic("min=.(int)")
	}
	return len(v.Interface().(string)) <= l
}

// IsGte gte定义
func IsGte(f reflect.StructField, v reflect.Value, param string) bool {
	if f.Type.Kind() != reflect.Int {
		panic("IsMax only supports int type")
	}
	max, err := strconv.Atoi(param)
	if err != nil {
		panic("gte=.(int)")
	}
	return v.Interface().(int) >= max
}

// IsLte lte定义
func isLte(f reflect.StructField, v reflect.Value, param string) bool {
	if f.Type.Kind() != reflect.Int {
		panic("IsMax only supports int type")
	}
	min, err := strconv.Atoi(param)
	if err != nil {
		panic("lte=.(int)")
	}
	return v.Interface().(int) <= min
}
