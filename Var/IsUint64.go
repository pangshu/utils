package Var

// IsUint64 判断变量中否为uint64类型.
func (conv *Var) IsUint64(value any) bool {
	switch value.(type) {
	case uint64:
		return true
	default:
		return false
	}
}
