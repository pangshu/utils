package Var

import (
	"fmt"
	"math"
	"strconv"
)

// ToUint 强制将变量转换为uint类型.
func (conv *Var) ToUint(in any) (uint, error) {
	i := indirect(in)
	switch s := i.(type) {
	case int:
		if s < 0 {
			return uint(s), fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
		} else {
			return uint(s), nil
		}
	case int8:
		if s < 0 {
			return uint(s), fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
		} else {
			return uint(s), nil
		}
	case int16:
		if s < 0 {
			return uint(s), fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
		} else {
			return uint(s), nil
		}
	case int32:
		if s < 0 {
			return uint(s), fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
		} else {
			return uint(s), nil
		}
	case int64:
		if s >= 0 {
			return uint(s), nil
		} else {
			return uint(s), fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
		}
	case uint:
		return s, nil
	case uint8:
		return uint(s), nil
	case uint16:
		return uint(s), nil
	case uint32:
		return uint(s), nil
	case uint64:
		if s <= math.MaxUint {
			return uint(s), nil
		} else {
			return uint(s), fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
		}
	case float32:
		if s <= math.MaxUint {
			return uint(s), nil
		} else {
			return uint(s), fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
		}
	case float64:
		if s <= math.MaxUint {
			return uint(s), nil
		} else {
			return uint(s), fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
		}
	case []byte:
		v, err := strconv.ParseUint(string(s), 0, 0)
		if err == nil {
			return uint(v), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case string:
		v, err := strconv.ParseUint(s, 0, 0)
		if err == nil {
			return uint(v), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
		}
	case bool:
		if s {
			return 1, nil
		} else {
			return 0, nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
	}
}
