package Var

// DefaultUint8 强制将变量转换为uint8类型，并设置返回默认值.
func (conv *Var) DefaultUint8(i any, def ...uint8) uint8 {
	if res, err := conv.ToUint8(i); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return 0
		} else {
			return def[0]
		}
	}
}
