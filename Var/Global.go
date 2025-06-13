package Var

import "reflect"

//type Var byte

type Var struct {
	value any
}

//type Var byte

func indirect(in any) any {
	if in == nil {
		return nil
	}
	if t := reflect.TypeOf(in); t.Kind() != reflect.Ptr {
		// 非指针类型，返回原值
		return in
	}
	v := reflect.ValueOf(in)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

var Variable = new(Var)

// // New creates and returns a new Var with given `value`.
// // The optional parameter `safe` specifies whether Var is used in concurrent-safety,
// // which is false in default.
func New(value any) *Var {
	return &Var{
		value: value,
	}
}
