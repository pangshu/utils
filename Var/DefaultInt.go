package Var

// DefaultInt 强制将变量转换为int类型，并设置返回默认值.
func (conv *Var) DefaultInt(i any, def ...int) int {
	if res, err := conv.ToInt(i); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return 0
		} else {
			return def[0]
		}
	}
}
