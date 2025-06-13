package Var

// DefaultString 强制将变量转换为string类型，并设置返回默认值.
func (conv *Var) DefaultString(value any, def ...string) string {
	if res, err := conv.ToString(value); err == nil {
		return res
	} else {
		if len(def) == 0 {
			return ""
		} else {
			return def[0]
		}
	}
}
