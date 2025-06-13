package Var

// DefaultUint32 强制将变量转换为uint32类型，并设置返回默认值.
func (conv *Var) DefaultUint32(i any, def ...uint32) uint32 {
	if res, err := conv.ToUint32(i); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return 0
		} else {
			return def[0]
		}
	}
}
