package Var

// IsString 判断变量中否为string类型.
func (conv *Var) IsString(value any) bool {
	switch value.(type) {
	case string:
		return true
	default:
		return false
	}
}
