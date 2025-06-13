package Var

// IsBool 判断变量中否为bool类型.
func (conv *Var) IsBool(value any) bool {
	switch value.(type) {
	case bool:
		return true
	default:
		return false
	}
}
