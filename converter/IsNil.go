package converter

import "reflect"

func (c *Converter[T]) IsNil(v T) bool {
	if nil == v {
		return true
	}
	return reflect.ValueOf(v).IsNil()
}
