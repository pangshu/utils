package Var

import (
	"fmt"
	"html/template"
	"reflect"
	"testing"
)

// go test -v -run TestToStr -bench=BenchmarkToStr -count=5 Var/*
//func BenchmarkToStr(t *testing.B) {
//	t.ResetTimer()
//	var toolVar Var
//	type Key struct {
//		k string
//	}
//	key := &Key{"foo"}
//
//	type Investment struct {
//		Price  float64
//		Symbol string
//		Rating int64
//	}
//	inv := Investment{Price: 534.432, Symbol: "GBG", Rating: 4}
//
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(8)},
//		{int8(8)},
//		{int16(8)},
//		{int32(8)},
//		{int64(8)},
//		{uint(8)},
//		{uint8(8)},
//		{uint16(8)},
//		{uint32(8)},
//		{uint64(8)},
//		{float32(8.31)},
//		{float64(8.31)},
//		{true},
//		{false},
//		{nil},
//		{[]byte("one time")},
//		{"one more time"},
//		{template.HTML("one time")},
//		{template.URL("http://somehost.foo")},
//		{template.JS("(1+2)")},
//		{template.CSS("a")},
//		{template.HTMLAttr("a")},
//		{func() error {return nil}},
//		// errors
//		{testing.T{}},
//		{key},
//		{&inv.Price},
//		{&inv.Symbol},
//		{&inv.Rating},
//		{inv},
//	}
//	for i:=0; i< t.N; i++ {
//		for _, test := range tests {
//			_,_ = toolVar.ToString(test.input)
//		}
//
//	}
//}

// /////////////////////////////////// 测试 ToStr ///////////////////////////////////
// 测试命令: go test -v -run TestToStr Var/*
func TestToString(t *testing.T) {
	var toolVar Var
	type Key struct {
		k string
	}
	key := &Key{"foo"}

	type Investment struct {
		Price  float64
		Symbol string
		Rating int64
	}
	inv := Investment{Price: 534.432, Symbol: "GBG", Rating: 4}

	tests := []struct {
		input interface{}
	}{
		{int(8)},
		{int8(8)},
		{int16(8)},
		{int32(8)},
		{int64(8)},
		{uint(8)},
		{uint8(8)},
		{uint16(8)},
		{uint32(8)},
		{uint64(8)},
		{float32(8.31)},
		{float64(8.31)},
		{true},
		{false},
		{nil},
		{[]byte("one time")},
		{"one more time"},
		{template.HTML("one time")},
		{template.URL("http://somehost.foo")},
		{template.JS("(1+2)")},
		{template.CSS("a")},
		{template.HTMLAttr("a")},
		{func() error { return nil }},
		// errors
		{testing.T{}},
		{key},
		{&inv.Price},
		{&inv.Symbol},
		{&inv.Rating},
		{inv},
	}

	for _, test := range tests {
		v, err := toolVar.ToString(test.input)
		b := reflect.ValueOf(test.input)
		fmt.Println(test.input, " >>>>>>>>>>>> ", b.Kind(), " >>>>>>>>>>> ", v, ">>>>>>>>>>>>>>>>>>", err)
	}
}

// go test -v -run TestToStr -bench=BenchmarkToStr -count=5 Var/*
func BenchmarkToStr(t *testing.B) {
	t.ResetTimer()
	var toolVar Var
	type Key struct {
		k string
	}
	key := &Key{"foo"}

	type Investment struct {
		Price  float64
		Symbol string
		Rating int64
	}
	inv := Investment{Price: 534.432, Symbol: "GBG", Rating: 4}

	tests := []struct {
		input interface{}
	}{
		{int(8)},
		{int8(8)},
		{int16(8)},
		{int32(8)},
		{int64(8)},
		{uint(8)},
		{uint8(8)},
		{uint16(8)},
		{uint32(8)},
		{uint64(8)},
		{float32(8.31)},
		{float64(8.31)},
		{true},
		{false},
		{nil},
		{[]byte("one time")},
		{"one more time"},
		{template.HTML("one time")},
		{template.URL("http://somehost.foo")},
		{template.JS("(1+2)")},
		{template.CSS("a")},
		{template.HTMLAttr("a")},
		{func() error { return nil }},
		// errors
		{testing.T{}},
		{key},
		{&inv.Price},
		{&inv.Symbol},
		{&inv.Rating},
		{inv},
	}
	for i := 0; i < t.N; i++ {
		for _, test := range tests {
			_, _ = toolVar.ToString(test.input)
		}

	}
}

// /////////////////////////////////// 测试 ToInt ///////////////////////////////////
// 测试命令: go test -v -run TestToInt Var/*
func TestToInt(t *testing.T) {
	var toolVar Var
	type Investment struct {
		Price  float64
		Symbol string
		Rating int64
	}
	inv := Investment{Price: 534.432, Symbol: "GBG", Rating: 4}
	tests := []struct {
		input interface{}
	}{
		{int(9223372036854775807)},
		{int8(127)},
		{int16(32767)},
		{int32(2147483647)},
		{int64(9223372036854775807)},
		{uint(18446744073709551615)},
		{uint8(255)},
		{uint16(65535)},
		{uint32(4294967295)},
		{uint64(18446744073709551615)},
		{float32(8.31)},
		{float64(8.31)},
		{true},
		{false},
		{"9223372036854775807"},
		{nil},
		// errors
		{"test"},
		{testing.T{}},
		{&inv.Price},
		{&inv.Symbol},
		{&inv.Rating},
	}

	for _, test := range tests {
		//errmsg := fmt.Sprintf("i = %d", i) // assert helper message
		//fmt.Println(errmsg)
		v, err := toolVar.ToInt(test.input)
		b := reflect.ValueOf(test.input)
		fmt.Println(test.input, " >>>>>>>>>>>> ", b.Kind(), " >>>>>>>>>>> ", v, ">>>>>>>>>>>>>>>>>>", err)
	}
}

// go test -v -run TestToInt -bench=BenchmarkToInt -count=5 Var/*
func BenchmarkToInt(t *testing.B) {
	t.ResetTimer()
	var toolVar Var
	type Investment struct {
		Price  float64
		Symbol string
		Rating int64
	}
	inv := Investment{Price: 534.432, Symbol: "GBG", Rating: 4}
	tests := []struct {
		input interface{}
	}{
		{int(9223372036854775807)},
		{int8(127)},
		{int16(32767)},
		{int32(2147483647)},
		{int64(9223372036854775807)},
		{uint(18446744073709551615)},
		{uint8(255)},
		{uint16(65535)},
		{uint32(4294967295)},
		{uint64(18446744073709551615)},
		{float32(8.31)},
		{float64(8.31)},
		{true},
		{false},
		{"9223372036854775807"},
		{nil},
		// errors
		{"test"},
		{testing.T{}},
		{&inv.Price},
		{&inv.Symbol},
		{&inv.Rating},
	}
	for i := 0; i < t.N; i++ {
		for _, test := range tests {
			_, _ = toolVar.ToInt(test.input)
		}

	}
}

/////////////////////////////////////// 测试 ToInt8 ///////////////////////////////////
//// 测试命令: go test -v -run TestToInt8 Var/*
//func TestToInt8(t *testing.T) {
//	var toolVar Var
//	type Investment struct {
//		Price  float64
//		Symbol string
//		Rating int64
//	}
//	inv := Investment{Price: 534.432, Symbol: "GBG", Rating: 4}
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(8.31)},
//		{float64(8.31)},
//		{true},
//		{false},
//		{"9223372036854775807"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//		{&inv.Price},
//		{&inv.Symbol},
//		{&inv.Rating},
//	}
//
//	for _, test := range tests {
//		//errmsg := fmt.Sprintf("i = %d", i) // assert helper message
//		//fmt.Println(errmsg)
//		v,err := toolVar.ToInt8(test.input)
//		b := reflect.ValueOf(test.input)
//		fmt.Println(test.input , " >>>>>>>>>>>> " , b.Kind() , " >>>>>>>>>>> ", v, ">>>>>>>>>>>>>>>>>>", err)
//	}
//
//}
//
//// go test -v -run TestToInt8 -bench=BenchmarkToInt8 -count=5 Var/*
//func BenchmarkToInt8(t *testing.B) {
//	t.ResetTimer()
//	var toolVar Var
//	type Investment struct {
//		Price  float64
//		Symbol string
//		Rating int64
//	}
//	inv := Investment{Price: 534.432, Symbol: "GBG", Rating: 4}
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(8.31)},
//		{float64(8.31)},
//		{true},
//		{false},
//		{"9223372036854775807"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//		{&inv.Price},
//		{&inv.Symbol},
//		{&inv.Rating},
//	}
//	for i:=0; i< t.N; i++ {
//		for _, test := range tests {
//			_,_ = toolVar.ToInt8(test.input)
//		}
//
//	}
//}

/////////////////////////////////////// 测试 ToInt16 ///////////////////////////////////
//// 测试命令: go test -v -run TestToInt16 Var/*
//func TestToInt16(t *testing.T) {
//	var toolVar Var
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(88888888.31)},
//		{float64(88888888.31)},
//		{true},
//		{false},
//		{"9223372036854775807"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//	}
//
//	for _, test := range tests {
//		//errmsg := fmt.Sprintf("i = %d", i) // assert helper message
//		//fmt.Println(errmsg)
//		v,err := toolVar.ToInt16(test.input)
//		b := reflect.ValueOf(test.input)
//		fmt.Println(test.input , " >>>>>>>>>>>> " , b.Kind() , " >>>>>>>>>>> ", v, ">>>>>>>>>>>>>>>>>>", err)
//	}
//}
//
//// go test -v -run TestToInt16 -bench=BenchmarkToInt16 -count=5 Var/*
//func BenchmarkToInt16(t *testing.B) {
//	t.ResetTimer()
//	var toolVar Var
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(8.31)},
//		{float64(8.31)},
//		{true},
//		{false},
//		{"8"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//	}
//	for i:=0; i< t.N; i++ {
//		for _, test := range tests {
//			_,_ = toolVar.ToInt16(test.input)
//		}
//
//	}
//}
//
/////////////////////////////////////// 测试 ToInt32 ///////////////////////////////////
//// 测试命令: go test -v -run TestToInt32 Var/*
//func TestToInt32(t *testing.T) {
//	var toolVar Var
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(88888888.31)},
//		{float64(88888888.31)},
//		{true},
//		{false},
//		{"9223372036854775807"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//	}
//
//	for _, test := range tests {
//		//errmsg := fmt.Sprintf("i = %d", i) // assert helper message
//		//fmt.Println(errmsg)
//		v,err := toolVar.ToInt32(test.input)
//		b := reflect.ValueOf(test.input)
//		fmt.Println(test.input , " >>>>>>>>>>>> " , b.Kind() , " >>>>>>>>>>> ", v, ">>>>>>>>>>>>>>>>>>", err)
//	}
//}
//
//// go test -v -run TestToInt32 -bench=BenchmarkToInt32 -count=5 Var/*
//func BenchmarkToInt32(t *testing.B) {
//	t.ResetTimer()
//	var toolVar Var
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(8.31)},
//		{float64(8.31)},
//		{true},
//		{false},
//		{"8"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//	}
//	for i:=0; i< t.N; i++ {
//		for _, test := range tests {
//			_,_ = toolVar.ToInt32(test.input)
//		}
//
//	}
//}
//
/////////////////////////////////////// 测试 ToInt64 ///////////////////////////////////
//// 测试命令: go test -v -run TestToInt64 Var/*
//func TestToInt64(t *testing.T) {
//	var toolVar Var
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(88888888.31)},
//		{float64(88888888.31)},
//		{true},
//		{false},
//		{"9223372036854775807"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//	}
//
//	for _, test := range tests {
//		//errmsg := fmt.Sprintf("i = %d", i) // assert helper message
//		//fmt.Println(errmsg)
//		v,err := toolVar.ToInt64(test.input)
//		b := reflect.ValueOf(test.input)
//		fmt.Println(test.input , " >>>>>>>>>>>> " , b.Kind() , " >>>>>>>>>>> ", v, ">>>>>>>>>>>>>>>>>>", err)
//	}
//}
//
//// go test -v -run TestToInt64 -bench=BenchmarkToInt64 -count=5 Var/*
//func BenchmarkToInt64(t *testing.B) {
//	t.ResetTimer()
//	var toolVar Var
//	tests := []struct {
//		input  interface{}
//	}{
//		{int(9223372036854775807)},
//		{int8(127)},
//		{int16(32767)},
//		{int32(2147483647)},
//		{int64(9223372036854775807)},
//		{uint(18446744073709551615)},
//		{uint8(255)},
//		{uint16(65535)},
//		{uint32(4294967295)},
//		{uint64(18446744073709551615)},
//		{float32(8.31)},
//		{float64(8.31)},
//		{true},
//		{false},
//		{"8"},
//		{nil},
//		// errors
//		{"test"},
//		{testing.T{}},
//	}
//	for i:=0; i< t.N; i++ {
//		for _, test := range tests {
//			_,_ = toolVar.ToInt64(test.input)
//		}
//
//	}
//}

// TestIsBool 测试 IsBool 方法的各种情况
// go test -v -run TestIsBool -bench=BenchmarkToInt -count=5 Var/*
func TestIsBool(t *testing.T) {
	var conv Var

	res := conv.IsBool(nil)
	fmt.Println(res)
	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{
			name:     "True boolean",
			input:    true,
			expected: true,
		},
		{
			name:     "False boolean",
			input:    false,
			expected: true,
		},
		{
			name:     "Integer input",
			input:    123,
			expected: false,
		},
		{
			name:     "String input",
			input:    "true",
			expected: false,
		},
		{
			name:     "Nil input",
			input:    nil,
			expected: false,
		},
		{
			name:     "Struct input",
			input:    struct{}{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := conv.IsBool(tt.input)
			if result != tt.expected {
				t.Errorf("IsBool(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
