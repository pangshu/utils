package Var

// DefaultInt8 强制将变量转换为int8类型，并设置返回默认值.
func (conv *Var) DefaultInt8(i any, def ...int8) int8 {
	if res, err := conv.ToInt8(i); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return 0
		} else {
			return def[0]
		}
	}
}
