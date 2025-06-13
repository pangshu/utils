package Var

// DefaultFloat32 强制将变量转换为float32类型，并设置返回默认值.
func (conv *Var) DefaultFloat32(i any, def ...float32) float32 {
	if res, err := conv.ToFloat32(i); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return 0
		} else {
			return def[0]
		}
	}
}
