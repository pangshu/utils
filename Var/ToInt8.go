package Var

import (
	"fmt"
	"math"
	"strconv"
)

// ToInt8 强制将变量转换为int8类型.
func (conv *Var) ToInt8(in any) (int8, error) {
	i := indirect(in)
	switch s := i.(type) {
	case int:
		if s >= math.MinInt8 && s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case int8:
		return s, nil
	case int16:
		if s >= math.MinInt8 && s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case int32:
		if s >= math.MinInt8 && s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case int64:
		if s >= math.MinInt8 && s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case uint:
		if s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case uint8:
		if s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case uint16:
		if s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case uint32:
		if s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case uint64:
		if s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case float32:
		if s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case float64:
		if s <= math.MaxInt8 {
			return int8(s), nil
		} else {
			return int8(s), fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
		}
	case []byte:
		v, err := strconv.ParseInt(string(s), 0, 8)
		if err == nil {
			return int8(v), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case string:
		v, err := strconv.ParseInt(s, 0, 8)
		if err == nil {
			return int8(v), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
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
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
	}
}
