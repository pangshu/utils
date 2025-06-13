package Var

// DefaultUint16 强制将变量转换为uint16类型，并设置返回默认值.
func (conv *Var) DefaultUint16(i any, def ...uint16) uint16 {
	if res, err := conv.ToUint16(i); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return 0
		} else {
			return def[0]
		}
	}
}
