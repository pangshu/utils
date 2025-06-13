package Var

// DefaultInt16 强制将变量转换为int16类型，并设置返回默认值.
func (conv *Var) DefaultInt16(i any, def ...int16) int16 {
	if res, err := conv.ToInt16(i); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return 0
		} else {
			return def[0]
		}
	}
}
