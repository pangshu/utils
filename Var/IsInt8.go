package Var

// IsInt8 判断变量中否为int8类型.
func (conv *Var) IsInt8(value any) bool {
	switch value.(type) {
	case int8:
		return true
	default:
		return false
	}
}
