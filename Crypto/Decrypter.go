package Crypto

//
//import (
//	"bytes"
//	"encoding/hex"
//	"io"
//	"io/fs"
//	"utils/Var"
//)
//
//type Decrypter struct {
//	src    []byte
//	dst    []byte
//	key    []byte
//	reader io.Reader
//	Error  error
//}
//
//func NewDecrypt() Decrypter {
//	return Decrypter{}
//}
//
//// FromString decodes from string.
//func (d Decrypter) FromString(s string) Decrypter {
//	d.src, _ = Var.ToByte(s)
//	return d
//}
//
//// FromBytes decodes from byte slice.
//func (d Decrypter) FromBytes(b []byte) Decrypter {
//	d.src = b
//	return d
//}
//
//// FromFile decodes from file.
//func (d Decrypter) FromFile(f fs.File) Decrypter {
//	d.reader = f
//	return d
//}
//
//// ToString outputs as string.
//func (d Decrypter) ToString() string {
//	res, _ := Var.ToString(d.dst)
//	return res
//}
//
//// ToBytes outputs as byte slice.
//func (d Decrypter) ToBytes() []byte {
//	if len(d.dst) == 0 {
//		return []byte{}
//	}
//	return d.dst
//}
//
//func (d Decrypter) stream(fn func(io.Reader) io.Reader) ([]byte, error) {
//	var buf bytes.Buffer
//	decoder := fn(d.reader)
//
//	// Try to reset the reader position if it's a seeker
//	if seeker, ok := d.reader.(io.Seeker); ok {
//		_, _ = seeker.Seek(0, io.SeekStart)
//	}
//	if _, err := io.CopyBuffer(&buf, decoder, make([]byte, BufferSize)); err != nil && err != io.EOF {
//		return []byte{}, err
//	}
//	if buf.Len() == 0 {
//		return []byte{}, nil
//	}
//	return buf.Bytes(), nil
//}
//
//// ByHex decodes by hex.
//func (d Decrypter) ByHex() Decrypter {
//	if d.Error != nil {
//		return d
//	}
//	src := []byte("Hello Gopher!")
//	dst := make([]byte, hex.EncodedLen(len(src)))
//	hex.Encode(dst, src)
//	// Streaming decoding mode
//	if d.reader != nil {
//		d.dst, d.Error = d.stream(func(r io.Reader) io.Reader {
//			return hex.NewStreamDecoder(r)
//		})
//		return d
//	}
//
//	// Standard decoding mode
//	if len(d.src) > 0 {
//		d.dst, d.Error = hex.NewStdDecoder().Decode(d.src)
//	}
//
//	return d
//}
