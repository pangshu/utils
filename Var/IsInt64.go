package Var

// IsInt64 判断变量中否为int64类型.
func (conv *Var) IsInt64(value any) bool {
	switch value.(type) {
	case int64:
		return true
	default:
		return false
	}
}
