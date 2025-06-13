package Var

import (
	"fmt"
	"math"
	"strconv"
)

// ToInt32 强制将变量转换为int32类型.
func (conv *Var) ToInt32(in any) (int32, error) {
	i := indirect(in)
	//处理类型转换
	switch s := i.(type) {
	case int:
		if s >= math.MinInt32 && s <= math.MaxInt32 {
			return int32(s), nil
		} else {
			return int32(s), fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
		}
	case int8:
		return int32(s), nil
	case int16:
		return int32(s), nil
	case int32:
		return s, nil
	case int64:
		if s >= math.MinInt32 && s <= math.MaxInt32 {
			return int32(s), nil
		} else {
			return int32(s), fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
		}
	case uint:
		if s <= math.MaxInt32 {
			return int32(s), nil
		} else {
			return int32(s), fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
		}
	case uint8:
		return int32(s), nil
	case uint16:
		return int32(s), nil
	case uint32:
		if s <= math.MaxInt32 {
			return int32(s), nil
		} else {
			return int32(s), fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
		}
	case uint64:
		if s <= math.MaxInt32 {
			return int32(s), nil
		} else {
			return int32(s), fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
		}
	case float32:
		if s <= math.MaxInt32 {
			return int32(s), nil
		} else {
			return int32(s), fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
		}
	case float64:
		if s <= math.MaxInt32 {
			return int32(s), nil
		} else {
			return int32(s), fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
		}
	case []byte:
		vv, err := strconv.ParseInt(string(s), 0, 32)
		if err == nil {
			return int32(vv), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case string:
		vv, err := strconv.ParseInt(s, 0, 32)
		if err == nil {
			return int32(vv), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
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
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
	}
}
