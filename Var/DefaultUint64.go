package Var

// DefaultUint64 强制将变量转换为uint64类型，并设置返回默认值.
func (conv *Var) DefaultUint64(i any, def ...uint64) uint64 {
	if res, err := conv.ToUint64(i); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return 0
		} else {
			return def[0]
		}
	}
}
