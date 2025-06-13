package Var

import (
	"fmt"
	"math"
	"strconv"
)

// ToUint32 强制将变量转换为uint32类型.
func (conv *Var) ToUint32(in any) (uint32, error) {
	i := indirect(in)
	switch s := i.(type) {
	case int:
		if s >= 0 && s <= math.MaxUint32 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case int8:
		if s >= 0 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case int16:
		if s >= 0 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case int32:
		if s >= 0 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case int64:
		if s >= 0 && s <= math.MaxUint32 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case uint:
		if s <= math.MaxUint32 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case uint8:
		return uint32(s), nil
	case uint16:
		return uint32(s), nil
	case uint32:
		return s, nil
	case uint64:
		if s <= math.MaxUint32 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case float32:
		if s <= math.MaxUint32 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case float64:
		if s <= math.MaxUint32 {
			return uint32(s), nil
		} else {
			return uint32(s), fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
		}
	case []byte:
		v, err := strconv.ParseUint(string(s), 0, 32)
		if err == nil {
			return uint32(v), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case string:
		v, err := strconv.ParseUint(s, 0, 32)
		if err == nil {
			return uint32(v), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
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
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
	}
}
