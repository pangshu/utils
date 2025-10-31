package Crypto

import (
	"bytes"
	"crypto/cipher"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"os"
	"reflect"
	"strconv"
)

type Encrypter struct {
	src       []byte
	dst       []byte
	key       []byte
	BlockMode cipher.BlockMode
	reader    io.Reader
	Error     error
}

func NewEncrypt() Encrypter {
	return Encrypter{}
}

func (e *Encrypter) Content(in any) *Encrypter {
	i := e.indirect(in)
	switch s := i.(type) {
	case string:
		e.src = []byte(s)
		return e
	case []byte:
		e.src = s
		return e
	case *os.File:
		var buf bytes.Buffer
		decoder := fn(i.reader)

		// Try to reset the reader position if it's a seeker
		if seeker, ok := d.reader.(io.Seeker); ok {
			seeker.Seek(0, io.SeekStart)
		}
		if _, err := io.CopyBuffer(&buf, decoder, make([]byte, BufferSize)); err != nil && err != io.EOF {
			return []byte{}, err
		}
		if buf.Len() == 0 {
			return []byte{}, nil
		}
		return buf.Bytes(), nil
		e.src =
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to string", i, i)
	}
}

// // FromString encodes from string.
//
//	func (e Encrypter) FromString(s string) Encrypter {
//		e.src, _ = Var.ToByte(s)
//		return e
//	}
//
// // FromBytes encodes from byte slice.
//
//	func (e Encrypter) FromBytes(b []byte) Encrypter {
//		e.src = b
//		return e
//	}
//
// // FromFile encodes from file.
//
//	func (e Encrypter) FromFile(f fs.File) Encrypter {
//		e.reader = f
//		return e
//	}
//
// // ToString outputs as string.
//
//	func (e Encrypter) ToString() string {
//		res, _ := Var.ToString(e.dst)
//		return res
//	}
//
// // ToBytes outputs as byte slice.
//
//	func (e Encrypter) ToBytes() []byte {
//		if len(e.dst) == 0 {
//			return []byte{}
//		}
//		return e.dst
//	}
//
// //func (e Encrypter) ToBase64String() string {
// //	return coding.NewEncoder().FromBytes(e.dst).ByBase64().ToString()
// //}
// //
// //// ToBase64Bytes outputs as base64 byte slice.
// //func (e Encrypter) ToBase64Bytes() []byte {
// //	return coding.NewEncoder().FromBytes(e.dst).ByBase64().ToBytes()
// //}
// //
// //// ToHexString outputs as hex string.
// //func (e Encrypter) ToHexString() string {
// //	return coding.NewEncoder().FromBytes(e.dst).ByHex().ToString()
// //}
// //
// //// ToHexBytes outputs as hex byte slice.
// //func (e Encrypter) ToHexBytes() []byte {
// //	return coding.NewEncoder().FromBytes(e.dst).ByHex().ToBytes()
// //}
//
//	func (e Encrypter) stream(fn func(io.Writer) io.WriteCloser) ([]byte, error) {
//		var buf bytes.Buffer
//		encoder := fn(&buf)
//
//		// Try to reset the reader position if it's a seeker
//		if seeker, ok := e.reader.(io.Seeker); ok {
//			_, _ = seeker.Seek(0, io.SeekStart)
//		}
//		if _, err := io.CopyBuffer(encoder, e.reader, make([]byte, BufferSize)); err != nil && err != io.EOF {
//			_ = encoder.Close()
//			return []byte{}, err
//		}
//		if err := encoder.Close(); err != nil {
//			return []byte{}, err
//		}
//		if buf.Len() == 0 {
//			return []byte{}, nil
//		}
//		return buf.Bytes(), nil
//	}
//
// // // ByHex encodes by hex.
// //
// //	func (e Encrypter) ByHex() Encrypter {
// //		if e.Error != nil {
// //			return e
// //		}
// //
// //		// Streaming encoding mode
// //		if e.reader != nil {
// //			e.dst, e.Error = e.stream(func(w io.Writer) io.WriteCloser {
// //				return hex.NewStreamEncoder(w)
// //			})
// //			return e
// //		}
// //
// //		// Standard encoding mode
// //		if len(e.src) > 0 {
// //			e.dst = hex.NewStdEncoder().Encode(e.src)
// //		}
// //
// //		return e
// //	}
func (e *Encrypter) indirect(in any) any {
	if in == nil {
		return nil
	}
	if t := reflect.TypeOf(in); t.Kind() != reflect.Ptr {
		// 非指针类型，返回原值
		return in
	}
	v := reflect.ValueOf(in)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}
