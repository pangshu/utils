package Var

// IsInt16 判断变量中否为int16类型.
func (conv *Var) IsInt16(value any) bool {
	switch value.(type) {
	case int16:
		return true
	default:
		return false
	}
}
