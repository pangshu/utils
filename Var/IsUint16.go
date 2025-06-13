package Var

// IsUint16 判断变量中否为uint16类型.
func (conv *Var) IsUint16(value any) bool {
	switch value.(type) {
	case uint16:
		return true
	default:
		return false
	}
}
