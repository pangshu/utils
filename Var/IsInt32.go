package Var

// IsInt32 判断变量中否为int32类型.
func (conv *Var) IsInt32(value any) bool {
	switch value.(type) {
	case int32:
		return true
	default:
		return false
	}
}
