package Var

// DefaultInt32 强制将变量转换为int32类型，并设置返回默认值.
func (conv *Var) DefaultInt32(i any, def ...int32) int32 {
	if res, err := conv.ToInt32(i); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return 0
		} else {
			return def[0]
		}
	}
}
