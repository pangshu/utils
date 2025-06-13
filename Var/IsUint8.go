package Var

// IsUint8 判断变量中否为uint8类型.
func (conv *Var) IsUint8(value any) bool {
	switch value.(type) {
	case uint8:
		return true
	default:
		return false
	}
}
