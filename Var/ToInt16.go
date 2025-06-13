package Var

import (
	"fmt"
	"math"
	"strconv"
)

// ToInt16 强制将变量转换为int16类型.
func (conv *Var) ToInt16(in any) (int16, error) {
	i := indirect(in)
	//处理类型转换
	switch s := i.(type) {
	case int:
		if s >= math.MinInt16 && s <= math.MaxInt16 {
			return int16(s), nil
		} else {
			return int16(s), fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
		}
	case int8:
		return int16(s), nil
	case int16:
		return s, nil
	case int32:
		if s >= math.MinInt16 && s <= math.MaxInt16 {
			return int16(s), nil
		} else {
			return int16(s), fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
		}
	case int64:
		if s >= math.MinInt16 && s <= math.MaxInt16 {
			return int16(s), nil
		} else {
			return int16(s), fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
		}
	case uint:
		if s <= math.MaxInt16 {
			return int16(s), nil
		} else {
			return int16(s), fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
		}
	case uint8:
		return int16(s), nil
	case uint16:
		if s <= math.MaxInt16 {
			return int16(s), nil
		} else {
			return int16(s), fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
		}
	case uint32:
		if s <= math.MaxInt16 {
			return int16(s), nil
		} else {
			return int16(s), fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
		}
	case uint64:
		if s <= math.MaxInt16 {
			return int16(s), nil
		} else {
			return int16(s), fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
		}
	case float32:
		if s <= math.MaxInt16 {
			return int16(s), nil
		} else {
			return int16(s), fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
		}
	case float64:
		if s <= math.MaxInt16 {
			return int16(s), nil
		} else {
			return int16(s), fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
		}
	case []byte:
		v, err := strconv.ParseInt(string(s), 0, 16)
		if err == nil {
			return int16(v), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case string:
		vv, err := strconv.ParseInt(s, 0, 16)
		if err == nil {
			return int16(vv), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
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
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
	}
}
