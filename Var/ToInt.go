package Var

import (
	"fmt"
	"math"
	"strconv"
)

// ToInt 强制将变量转换为int类型.
func (conv *Var) ToInt(in any) (int, error) {
	i := indirect(in)
	//处理类型转换
	switch s := i.(type) {
	case int:
		return s, nil
	case int8:
		return int(s), nil
	case int16:
		return int(s), nil
	case int32:
		return int(s), nil
	case int64:
		if s >= math.MinInt && s <= math.MaxInt {
			return int(s), nil
		} else {
			return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case uint:
		if s <= math.MaxInt {
			return int(s), nil
		} else {
			return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case uint8:
		return int(s), nil
	case uint16:
		return int(s), nil
	case uint32:
		intSize := 32 << (^uint(0) >> 63) // 32 or 64
		if math.MaxUint32 >= (1<<(intSize-1) - 1) {
			if s <= (1<<(intSize-1) - 1) {
				return int(s), nil
			} else {
				return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
			}
		} else {
			return int(s), nil
		}
	case uint64:
		if s <= math.MaxInt {
			return int(s), nil
		} else {
			return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case float32:
		if s <= math.MaxInt {
			return int(s), nil
		} else {
			return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case float64:
		if s <= math.MaxInt {
			return int(s), nil
		} else {
			return int(s), fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case []byte:
		v, err := strconv.ParseInt(string(s), 0, 0)
		if err == nil {
			return int(v), nil
		} else {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int(v), nil
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
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	}
}
