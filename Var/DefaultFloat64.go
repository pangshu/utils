package Var

// DefaultFloat64 强制将变量转换为float64类型，并设置返回默认值.
func (conv *Var) DefaultFloat64(i any, def ...float64) float64 {
	if res, err := conv.ToFloat64(i); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return 0
		} else {
			return def[0]
		}
	}
}
