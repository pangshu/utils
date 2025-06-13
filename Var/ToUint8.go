package Var

import (
	"fmt"
	"math"
	"strconv"
)

// ToUint8 强制将变量转换为uint8类型.
func (conv *Var) ToUint8(in any) (uint8, error) {
	i := indirect(in)
	switch s := i.(type) {
	case int:
		if s >= 0 && s <= math.MaxUint8 {
			return uint8(s), nil
		} else {
			return uint8(s), fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
		}
	case int8:
		if s >= 0 {
			return uint8(s), nil
		} else {
			return uint8(s), fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
		}
	case int16:
		if s >= 0 && s <= math.MaxUint8 {
			return uint8(s), nil
		} else {
			return uint8(s), fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
		}
	case int32:
		if s >= 0 && s <= math.MaxUint8 {
			return uint8(s), nil
		} else {
			return uint8(s), fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
		}
	case int64:
		if s >= 0 && s <= math.MaxUint8 {
			return uint8(s), nil
		} else {
			return uint8(s), fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
		}
	case uint:
		if s <= math.MaxUint8 {
			return uint8(s), nil
		} else {
			return uint8(s), fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
		}
	case uint8:
		return s, nil
	case uint16:
		if s <= math.MaxUint8 {
			return uint8(s), nil
		} else {
			return uint8(s), fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
		}
	case uint32:
		if s <= math.MaxUint8 {
			return uint8(s), nil
		} else {
			return uint8(s), fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
		}
	case uint64:
		if s <= math.MaxUint8 {
			return uint8(s), nil
		} else {
			return uint8(s), fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
		}
	case float32:
		if s <= math.MaxUint8 {
			return uint8(s), nil
		} else {
			return uint8(s), fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
		}
	case float64:
		if s <= math.MaxUint8 {
			return uint8(s), nil
		} else {
			return uint8(s), fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
		}
	case []byte:
		v, err := strconv.ParseUint(string(s), 0, 8)
		if err == nil {
			return uint8(v), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case string:
		v, err := strconv.ParseUint(s, 0, 8)
		if err == nil {
			return uint8(v), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
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
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
	}
}
