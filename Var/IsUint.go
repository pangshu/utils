package Var

// IsUint 判断变量中否为uint类型.
func (conv *Var) IsUint(value any) bool {
	switch value.(type) {
	case uint:
		return true
	default:
		return false
	}
}
