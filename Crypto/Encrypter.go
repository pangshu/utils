package Crypto

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"reflect"
	"strconv"
)

// NewEncrypter 创建新的加密器实例
func NewEncrypter() *Encrypter {
	return &Encrypter{
		//chunkSize: 64 * 1024, // 默认64KB的块大小
	}
}

// Content 设置要加密/解密的内容，支持任意类型
func (e *Encrypter) Content(content any) *Encrypter {
	if e.Error != nil {
		return e
	}

	if content == nil {
		e.Error = fmt.Errorf("content cannot be nil")
		return e
	}

	// 根据不同类型转换为字节数组
	switch v := content.(type) {
	case string:
		e.src = []byte(v)
	case []byte:
		e.src = v
	case int, int8, int16, int32, int64:
		e.src = []byte(fmt.Sprintf("%d", v))
	case uint, uint8, uint16, uint32, uint64:
		e.src = []byte(fmt.Sprintf("%d", v))
	case float32, float64:
		e.src = []byte(fmt.Sprintf("%f", v))
	case bool:
		e.src = []byte(strconv.FormatBool(v))
	default:
		// 对于其他类型，尝试JSON序列化
		jsonData, err := json.Marshal(v)
		if err != nil {
			e.Error = fmt.Errorf("unsupported content type: %s, and JSON marshal failed: %v",
				reflect.TypeOf(content).String(), err)
			return e
		}
		e.src = jsonData
	}

	return e
}

//
//// FromFile 从文件读取内容（支持大文件）
//func (e *Encrypter) ContentFile(filePath string) *Encrypter {
//	if e.Error != nil {
//		return e
//	}
//
//	e.filePath = filePath
//	file, err := os.Open(filePath)
//	if err != nil {
//		e.Error = fmt.Errorf("failed to open file: %v", err)
//		return e
//	}
//	defer file.Close()
//
//	// 对于大文件，我们不立即读取全部内容，而是保存reader引用
//	e.reader = file
//	return e
//}

// ToString 返回字符串结果
func (e *Encrypter) ToString() string {
	if e.Error != nil {
		return ""
	}
	if e.dst == nil {
		return ""
	}
	return string(e.dst)
}

// ToHex 返回十六进制字符串
func (e *Encrypter) ToHex() string {
	if e.Error != nil {
		return ""
	}
	if e.dst == nil {
		return ""
	}
	return hex.EncodeToString(e.dst)
}

// ToBase64 返回Base64编码字符串
func (e *Encrypter) ToBase64() string {
	if e.Error != nil {
		return ""
	}
	if e.dst == nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(e.dst)
}

// ToBytes 返回字节数组
func (e *Encrypter) ToBytes() []byte {
	if e.Error != nil {
		return nil
	}
	return e.dst
}

// GetError 获取错误信息
func (e *Encrypter) GetError() error {
	return e.Error
}

// hash 通用哈希计算函数（内存数据）
func (e *Encrypter) hash(h hash.Hash) []byte {
	h.Write(e.src)
	return h.Sum(nil)
}

//
//// hashStream 流式哈希计算（大文件）
//func (e *Encrypter) hashStream(h hash.Hash) *Encrypter {
//	buffer := make([]byte, e.chunkSize)
//
//	for {
//		n, err := e.reader.Read(buffer)
//		if err != nil && err != io.EOF {
//			e.Error = fmt.Errorf("failed to read file: %v", err)
//			return e
//		}
//
//		if n == 0 {
//			break
//		}
//
//		h.Write(buffer[:n])
//	}
//
//	e.dst = h.Sum(nil)
//	return e
//}
//
//// aesEncryptStream 流式AES加密
//func (e *Encrypter) aesEncryptStream() *Encrypter {
//	key := e.padKey(e.key, 32)
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		e.Error = fmt.Errorf("failed to create AES cipher: %v", err)
//		return e
//	}
//
//	iv := key[:aes.BlockSize]
//	stream := cipher.NewCFBEncrypter(block, iv)
//
//	buffer := make([]byte, e.chunkSize)
//	for {
//		n, err := e.reader.Read(buffer)
//		if err != nil && err != io.EOF {
//			e.Error = fmt.Errorf("failed to read file: %v", err)
//			return e
//		}
//
//		if n == 0 {
//			break
//		}
//
//		ciphertext := make([]byte, n)
//		stream.XORKeyStream(ciphertext, buffer[:n])
//
//		if _, err := e.writer.Write(ciphertext); err != nil {
//			e.Error = fmt.Errorf("failed to write encrypted data: %v", err)
//			return e
//		}
//	}
//
//	return e
//}
//
//// aesDecryptStream 流式AES解密
//func (e *Encrypter) aesDecryptStream() *Encrypter {
//	key := e.padKey(e.key, 32)
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		e.Error = fmt.Errorf("failed to create AES cipher: %v", err)
//		return e
//	}
//
//	iv := key[:aes.BlockSize]
//	stream := cipher.NewCFBDecrypter(block, iv)
//
//	buffer := make([]byte, e.chunkSize)
//	for {
//		n, err := e.reader.Read(buffer)
//		if err != nil && err != io.EOF {
//			e.Error = fmt.Errorf("failed to read file: %v", err)
//			return e
//		}
//
//		if n == 0 {
//			break
//		}
//
//		plaintext := make([]byte, n)
//		stream.XORKeyStream(plaintext, buffer[:n])
//
//		if _, err := e.writer.Write(plaintext); err != nil {
//			e.Error = fmt.Errorf("failed to write decrypted data: %v", err)
//			return e
//		}
//	}
//
//	return e
//}

// padKey 填充密钥到指定长度
func (e *Encrypter) padKey(key []byte, size int) []byte {
	if len(key) >= size {
		return key[:size]
	}

	padded := make([]byte, size)
	copy(padded, key)
	// 用0填充剩余部分
	for i := len(key); i < size; i++ {
		padded[i] = 0
	}
	return padded
}

// Base64Decode Base64解码
func (e *Encrypter) Base64Decode() *Encrypter {
	if e.Error != nil {
		return e
	}

	if len(e.src) == 0 {
		e.Error = fmt.Errorf("content is empty")
		return e
	}

	decoded, err := base64.StdEncoding.DecodeString(string(e.src))
	if err != nil {
		e.Error = fmt.Errorf("base64 decode failed: %v", err)
		return e
	}

	e.dst = decoded
	return e
}

// HexDecode 十六进制解码
func (e *Encrypter) HexDecode() *Encrypter {
	if e.Error != nil {
		return e
	}

	if len(e.src) == 0 {
		e.Error = fmt.Errorf("content is empty")
		return e
	}

	decoded, err := hex.DecodeString(string(e.src))
	if err != nil {
		e.Error = fmt.Errorf("hex decode failed: %v", err)
		return e
	}

	e.dst = decoded
	return e
}

//
//// GetFileHash 快速获取文件哈希（便捷方法）
//func GetFileHash(filePath string, algorithm string) (string, error) {
//	encrypter := NewEncrypter().FromFile(filePath)
//
//	switch algorithm {
//	case "md5":
//		encrypter.ByMd5()
//	case "sha256":
//		encrypter.BySha256()
//	default:
//		return "", fmt.Errorf("unsupported algorithm: %s", algorithm)
//	}
//
//	if encrypter.Error != nil {
//		return "", encrypter.Error
//	}
//
//	return encrypter.ToHex(), nil
//}

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
