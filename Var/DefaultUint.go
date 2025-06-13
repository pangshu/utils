package Var

// DefaultUint 强制将变量转换为uint类型，并设置返回默认值.
func (conv *Var) DefaultUint(i any, def ...uint) uint {
	if res, err := conv.ToUint(i); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return 0
		} else {
			return def[0]
		}
	}
}
