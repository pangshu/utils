package Var

// IsInt 判断变量中否为int类型.
func (conv *Var) IsInt(value any) bool {
	switch value.(type) {
	case int:
		return true
	default:
		return false
	}
}
