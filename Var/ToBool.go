package Var

import (
	"fmt"
	"strconv"
)

// ToBool 强制将变量转换为bool类型.
func (conv *Var) ToBool(in any) (bool, error) {
	i := indirect(in)
	switch s := i.(type) {
	case bool:
		return s, nil
	case int:
		return s != 0, nil
	case int8:
		return s != 0, nil
	case int32:
		return s != 0, nil
	case int64:
		return s != 0, nil
	case float32:
		return s != 0, nil
	case float64:
		return s != 0, nil
	case uint:
		return s != 0, nil
	case uint8:
		return s != 0, nil
	case uint32:
		return s != 0, nil
	case uint64:
		return s != 0, nil
	case []byte:
		return strconv.ParseBool(i.(string))
	case string:
		return strconv.ParseBool(i.(string))
	default:
		return false, fmt.Errorf("unable to cast %#v of type %T to bool", i, i)
	}
}
