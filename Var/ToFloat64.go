package Var

import (
	"fmt"
	"strconv"
)

// ToFloat64 强制将变量转换为整float64类型.
func (conv *Var) ToFloat64(in any) (float64, error) {
	i := indirect(in)
	switch s := i.(type) {
	case float64:
		return s, nil
	case float32:
		return float64(s), nil
	case int:
		return float64(s), nil
	case int64:
		return float64(s), nil
	case int32:
		return float64(s), nil
	case int16:
		return float64(s), nil
	case int8:
		return float64(s), nil
	case uint:
		return float64(s), nil
	case uint64:
		return float64(s), nil
	case uint32:
		return float64(s), nil
	case uint16:
		return float64(s), nil
	case uint8:
		return float64(s), nil
	case []byte:
		v, err := strconv.ParseFloat(string(s), 64)
		if err == nil {
			return v, nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case string:
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return v, nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to float64", i, i)
		}
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", i, i)
	}
}
