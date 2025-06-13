package Var

// DefaultInt64 强制将变量转换为int64类型，并设置返回默认值.
func (conv *Var) DefaultInt64(i any, def ...int64) int64 {
	if res, err := conv.ToInt64(i); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return 0
		} else {
			return def[0]
		}
	}
}
