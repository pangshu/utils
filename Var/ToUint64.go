package Var

import (
	"fmt"
	"math"
	"strconv"
)

// ToUint64 强制将变量转换为uint64类型.
func (conv *Var) ToUint64(in any) (uint64, error) {
	i := indirect(in)
	switch s := i.(type) {
	case int:
		if s >= 0 {
			return uint64(s), nil
		} else {
			return uint64(s), fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
		}
	case int8:
		if s >= 0 {
			return uint64(s), nil
		} else {
			return uint64(s), fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
		}
	case int16:
		if s >= 0 {
			return uint64(s), nil
		} else {
			return uint64(s), fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
		}
	case int32:
		if s >= 0 {
			return uint64(s), nil
		} else {
			return uint64(s), fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
		}
	case int64:
		if s >= 0 {
			return uint64(s), nil
		} else {
			return uint64(s), fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
		}
	case uint:
		return uint64(s), nil
	case uint8:
		return uint64(s), nil
	case uint16:
		return uint64(s), nil
	case uint32:
		return uint64(s), nil
	case uint64:
		return s, nil
	case float32:
		if s <= math.MaxUint64 {
			return uint64(s), nil
		} else {
			return uint64(s), fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
		}
	case float64:
		if s <= math.MaxUint64 {
			return uint64(s), nil
		} else {
			return uint64(s), fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
		}
	case []byte:
		v, err := strconv.ParseUint(string(s), 0, 64)
		if err == nil {
			return v, nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case string:
		v, err := strconv.ParseUint(s, 0, 64)
		if err == nil {
			return v, nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
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
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
	}
}
