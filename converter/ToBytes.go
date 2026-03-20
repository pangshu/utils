package converter

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"testing"
)

// ToBytes converts numeric types to Byte.
func (c *Converter[T]) ToBytes(v T) ([]byte, error) {
	if true == c.IsNil(v) {
		return nil, nil
	}

	switch val := any(v).(type) {
	case int, int8, int16, int32, int64:
		number := val.(int64)
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case uint, uint8, uint16, uint32, uint64:
		number := val.(uint64)
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case float32, float64:
		number := val.(float64)
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case bool:
		return []byte(strconv.FormatBool(val)), nil
	case []byte:
		return val, nil
	case string:
		return []byte(val), nil
	default:
		return nil, fmt.Errorf("unsupported type: %T for ToBytes", val)
	}
}

func TestConverter_ToBytes(t *testing.T) {
	// 测试 nil 值
	t.Run("NilValue", func(t *testing.T) {
		var c Converter[any]
		result, err := c.ToBytes(nil)
		if err != nil {
			t.Errorf("unexpected error for nil value: %v", err)
		}
		if result != nil {
			t.Errorf("expected nil result, got %v", result)
		}
	})
	// 测试整数类型
	t.Run("IntegerTypes", func(t *testing.T) {
		c := Converter[int64]{}
		testCases := []struct {
			name  string
			input int64
		}{
			{"Zero", 0},
			{"Positive", 12345},
			{"Negative", -12345},
			{"MaxInt64", math.MaxInt64},
			{"MinInt64", math.MinInt64},
		}
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result, err := c.ToBytes(tc.input)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				buf := new(bytes.Buffer)
				err = binary.Write(buf, binary.BigEndian, tc.input)
				if err != nil {
					t.Errorf("unexpected error creating expected bytes: %v", err)
				}
				if !bytes.Equal(result, buf.Bytes()) {
					t.Errorf("expected %v, got %v", buf.Bytes(), result)
				}
			})
		}
	})
	// 测试无符号整数类型
	t.Run("UnsignedIntegerTypes", func(t *testing.T) {
		c := Converter[uint64]{}
		testCases := []struct {
			name  string
			input uint64
		}{
			{"Zero", 0},
			{"Positive", 12345},
			{"MaxUint64", math.MaxUint64},
		}
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result, err := c.ToBytes(tc.input)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				buf := new(bytes.Buffer)
				err = binary.Write(buf, binary.BigEndian, tc.input)
				if err != nil {
					t.Errorf("unexpected error creating expected bytes: %v", err)
				}
				if !bytes.Equal(result, buf.Bytes()) {
					t.Errorf("expected %v, got %v", buf.Bytes(), result)
				}
			})
		}
	})
	// 测试浮点数类型
	t.Run("FloatTypes", func(t *testing.T) {
		c := Converter[float64]{}
		testCases := []struct {
			name  string
			input float64
		}{
			{"Zero", 0.0},
			{"Positive", 123.456},
			{"Negative", -123.456},
			{"MaxFloat64", math.MaxFloat64},
			{"SmallestNonzero", math.SmallestNonzeroFloat64},
		}
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result, err := c.ToBytes(tc.input)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				buf := new(bytes.Buffer)
				err = binary.Write(buf, binary.BigEndian, tc.input)
				if err != nil {
					t.Errorf("unexpected error creating expected bytes: %v", err)
				}
				if !bytes.Equal(result, buf.Bytes()) {
					t.Errorf("expected %v, got %v", buf.Bytes(), result)
				}
			})
		}
	})
	// 测试布尔类型
	t.Run("BoolType", func(t *testing.T) {
		c := Converter[bool]{}
		testCases := []struct {
			name  string
			input bool
			want  []byte
		}{
			{"True", true, []byte("true")},
			{"False", false, []byte("false")},
		}
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result, err := c.ToBytes(tc.input)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if !bytes.Equal(result, tc.want) {
					t.Errorf("expected %v, got %v", tc.want, result)
				}
			})
		}
	})
	// 测试字节切片
	t.Run("ByteSlice", func(t *testing.T) {
		c := Converter[[]byte]{}
		input := []byte{0x01, 0x02, 0x03}
		result, err := c.ToBytes(input)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !bytes.Equal(result, input) {
			t.Errorf("expected %v, got %v", input, result)
		}
	})
	// 测试字符串
	t.Run("String", func(t *testing.T) {
		c := Converter[string]{}
		input := "test string"
		result, err := c.ToBytes(input)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		expected := []byte(input)
		if !bytes.Equal(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
	// 测试不支持的类型
	t.Run("UnsupportedType", func(t *testing.T) {
		c := Converter[map[string]int]{}
		input := make(map[string]int)
		_, err := c.ToBytes(input)
		if err == nil {
			t.Error("expected error for unsupported type, got nil")
		}
		expectedErr := fmt.Sprintf("unsupported type: %T for ToBytes", input)
		if err.Error() != expectedErr {
			t.Errorf("expected error %q, got %q", expectedErr, err.Error())
		}
	})
}
