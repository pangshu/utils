package Var

// IsUint32 判断变量中否为uint32类型.
func (conv *Var) IsUint32(value any) bool {
	switch value.(type) {
	case uint32:
		return true
	default:
		return false
	}
}
