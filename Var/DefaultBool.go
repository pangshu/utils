package Var

// DefaultBool 强制将变量转换为bool类型，并设置返回默认值.
func (conv *Var) DefaultBool(i any, def ...bool) bool {
	if res, err := conv.ToBool(i); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return false
		} else {
			return def[0]
		}
	}
}
