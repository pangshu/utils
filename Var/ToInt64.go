package Var

import (
	"fmt"
	"math"
	"strconv"
)

// ToInt64 强制将变量转换为int64类型.
func (conv *Var) ToInt64(in any) (int64, error) {
	i := indirect(in)
	switch s := i.(type) {
	case int:
		if s >= math.MinInt64 && s <= math.MaxInt64 {
			return int64(s), nil
		} else {
			return int64(s), fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
		}
	case int8:
		return int64(s), nil
	case int16:
		return int64(s), nil
	case int32:
		return int64(s), nil

	case int64:
		return s, nil
	case uint:
		if s <= math.MaxInt64 {
			return int64(s), nil
		} else {
			return int64(s), fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
		}
	case uint8:
		return int64(s), nil
	case uint16:
		return int64(s), nil
	case uint32:
		return int64(s), nil
	case uint64:
		if s <= math.MaxInt64 {
			return int64(s), nil
		} else {
			return int64(s), fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
		}
	case float32:
		if s <= math.MaxInt64 {
			return int64(s), nil
		} else {
			return int64(s), fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
		}
	case float64:
		if s <= math.MaxInt64 {
			return int64(s), nil
		} else {
			return int64(s), fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
		}
	case []byte:
		vv, err := strconv.ParseInt(string(s), 0, 64)
		if err == nil {
			return vv, nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case string:
		vv, err := strconv.ParseInt(s, 0, 64)
		if err == nil {
			return vv, nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
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
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	}
}
