package Var

// DefaultByte 强制将变量转换为byte类型，并设置返回默认值.
func (conv *Var) DefaultByte(i any, def ...[]byte) []byte {
	if res, err := conv.ToByte(i); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return nil
		} else {
			return def[0]
		}
	}
}
