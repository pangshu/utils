package Var

import (
	"fmt"
	"html/template"
	"strconv"
)

// ToByte 强制将变量转换为byte类型.
func (conv *Var) ToByte(in any) ([]byte, error) {
	i := indirect(in)
	switch s := i.(type) {
	case string:
		return []byte(s), nil
	case bool:
		return []byte(strconv.FormatBool(s)), nil
	case int:
		return []byte(strconv.FormatInt(int64(s), 10)), nil
	case int8:
		return []byte(strconv.FormatInt(int64(s), 10)), nil
	case int16:
		return []byte(strconv.FormatInt(int64(s), 10)), nil
	case int32:
		return []byte(strconv.FormatInt(int64(s), 10)), nil
	case int64:
		return []byte(strconv.FormatInt(s, 10)), nil
	case uint:
		return []byte(strconv.FormatUint(uint64(s), 10)), nil
	case uint8:
		return []byte(strconv.FormatUint(uint64(s), 10)), nil
	case uint16:
		return []byte(strconv.FormatUint(uint64(s), 10)), nil
	case uint32:
		return []byte(strconv.FormatUint(uint64(s), 10)), nil
	case uint64:
		return []byte(strconv.FormatUint(s, 10)), nil
	case float32:
		return []byte(strconv.FormatFloat(float64(s), 'f', -1, 32)), nil
	case float64:
		return []byte(strconv.FormatFloat(s, 'f', -1, 64)), nil
	case []byte:
		return s, nil
	case template.HTML:
		return []byte(s), nil
	case template.URL:
		return []byte(s), nil
	case template.JS:
		return []byte(s), nil
	case template.CSS:
		return []byte(s), nil
	case template.HTMLAttr:
		return []byte(s), nil
	case nil:
		return nil, nil
	case fmt.Stringer:
		return []byte(s.String()), nil
	case error:
		return []byte(s.Error()), nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to string", i, i)
	}
}
