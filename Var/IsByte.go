package Var

// IsByte 判断变量中否为byte类型.
func (conv *Var) IsByte(value any) bool {
	switch value.(type) {
	case []byte:
		return true
	default:
		return false
	}
}
