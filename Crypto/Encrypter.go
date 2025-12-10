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

// Content 设置要加密/解密的内容，支持任意类型
func (c *Crypto) Content(content any) *Crypto {
	if c.Error != nil {
		return c
	}

	if content == nil {
		c.Error = fmt.Errorf("content cannot be nil")
		return c
	}

	c.src = c.anyToBytes(content)
	//// 根据不同类型转换为字节数组
	//switch v := content.(type) {
	//case string:
	//	c.src = []byte(v)
	//case []byte:
	//	c.src = v
	//case int, int8, int16, int32, int64:
	//	c.src = []byte(fmt.Sprintf("%d", v))
	//case uint, uint8, uint16, uint32, uint64:
	//	c.src = []byte(fmt.Sprintf("%d", v))
	//case float32, float64:
	//	c.src = []byte(fmt.Sprintf("%f", v))
	//case bool:
	//	c.src = []byte(strconv.FormatBool(v))
	//default:
	//	// 对于其他类型，尝试JSON序列化
	//	jsonData, err := json.Marshal(v)
	//	if err != nil {
	//		c.Error = fmt.Errorf("unsupported content type: %s, and JSON marshal failed: %v",
	//			reflect.TypeOf(content).String(), err)
	//		return c
	//	}
	//	c.src = jsonData
	//}

	return c
}

// anyToBytes 将任意类型转换为字节数组
func (c *Crypto) anyToBytes(v any) []byte {
	if v == nil {
		return nil
	}

	// 根据不同类型转换为字节数组
	switch val := v.(type) {
	case string:
		return []byte(val)
	case []byte:
		return val
	case int, int8, int16, int32, int64:
		return []byte(fmt.Sprintf("%d", val))
	case uint, uint8, uint16, uint32, uint64:
		return []byte(fmt.Sprintf("%d", val))
	case float32, float64:
		return []byte(fmt.Sprintf("%f", val))
	case bool:
		return []byte(strconv.FormatBool(val))
	default:
		// 对于其他类型，尝试JSON序列化
		jsonData, err := json.Marshal(val)
		if err != nil {
			return nil
		}
		return jsonData
	}

}

//// NewCrypto 创建新的加密器实例
//func NewCrypto() *Crypto {
//	return &Crypto{
//		//chunkSize: 64 * 1024, // 默认64KB的块大小
//	}
//}

//// Content 设置要加密/解密的内容，支持任意类型
//func (c *Crypto) Content(content any) *Crypto {
//	if c.Error != nil {
//		return c
//	}
//
//	if content == nil {
//		c.Error = fmt.Errorf("content cannot be nil")
//		return c
//	}
//
//	// 根据不同类型转换为字节数组
//	switch v := content.(type) {
//	case string:
//		c.src = []byte(v)
//	case []byte:
//		c.src = v
//	case int, int8, int16, int32, int64:
//		c.src = []byte(fmt.Sprintf("%d", v))
//	case uint, uint8, uint16, uint32, uint64:
//		c.src = []byte(fmt.Sprintf("%d", v))
//	case float32, float64:
//		c.src = []byte(fmt.Sprintf("%f", v))
//	case bool:
//		c.src = []byte(strconv.FormatBool(v))
//	default:
//		// 对于其他类型，尝试JSON序列化
//		jsonData, err := json.Marshal(v)
//		if err != nil {
//			c.Error = fmt.Errorf("unsupported content type: %s, and JSON marshal failed: %v",
//				reflect.TypeOf(content).String(), err)
//			return c
//		}
//		c.src = jsonData
//	}
//
//	return c
//}

//
//// FromFile 从文件读取内容（支持大文件）
//func (c *Crypto) ContentFile(filePath string) *Crypto {
//	if c.Error != nil {
//		return c
//	}
//
//	c.filePath = filePath
//	file, err := os.Open(filePath)
//	if err != nil {
//		c.Error = fmt.Errorf("failed to open file: %v", err)
//		return c
//	}
//	defer filc.Close()
//
//	// 对于大文件，我们不立即读取全部内容，而是保存reader引用
//	c.reader = file
//	return c
//}

// ToString 返回字符串结果
func (c *Crypto) ToString() string {
	if c.Error != nil {
		return ""
	}
	if c.dst == nil {
		return ""
	}
	return string(c.dst)
}

// ToHex 返回十六进制字符串
func (c *Crypto) ToHex() string {
	if c.Error != nil {
		return c.Error.Error()
	}
	if c.dst == nil {
		return ""
	}
	return hex.EncodeToString(c.dst)
}

// ToBase64 返回Base64编码字符串
func (c *Crypto) ToBase64() string {
	if c.Error != nil {
		return ""
	}
	if c.dst == nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(c.dst)
}

// ToBytes 返回字节数组
func (c *Crypto) ToBytes() []byte {
	if c.Error != nil {
		return nil
	}
	return c.dst
}

// GetError 获取错误信息
func (c *Crypto) GetError() error {
	return c.Error
}

// hash 通用哈希计算函数（内存数据）
func (c *Crypto) hash(h hash.Hash) []byte {
	h.Write(c.src)
	return h.Sum(nil)
}

//
//// hashStream 流式哈希计算（大文件）
//func (c *Crypto) hashStream(h hash.Hash) *Crypto {
//	buffer := make([]byte, c.chunkSize)
//
//	for {
//		n, err := c.reader.Read(buffer)
//		if err != nil && err != io.EOF {
//			c.Error = fmt.Errorf("failed to read file: %v", err)
//			return c
//		}
//
//		if n == 0 {
//			break
//		}
//
//		h.Write(buffer[:n])
//	}
//
//	c.dst = h.Sum(nil)
//	return c
//}
//
//// aesEncryptStream 流式AES加密
//func (c *Crypto) aesEncryptStream() *Crypto {
//	key := c.padKey(c.key, 32)
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		c.Error = fmt.Errorf("failed to create AES cipher: %v", err)
//		return c
//	}
//
//	iv := key[:aes.BlockSize]
//	stream := cipher.NewCFBCrypto(block, iv)
//
//	buffer := make([]byte, c.chunkSize)
//	for {
//		n, err := c.reader.Read(buffer)
//		if err != nil && err != io.EOF {
//			c.Error = fmt.Errorf("failed to read file: %v", err)
//			return c
//		}
//
//		if n == 0 {
//			break
//		}
//
//		ciphertext := make([]byte, n)
//		stream.XORKeyStream(ciphertext, buffer[:n])
//
//		if _, err := c.writer.Write(ciphertext); err != nil {
//			c.Error = fmt.Errorf("failed to write encrypted data: %v", err)
//			return c
//		}
//	}
//
//	return c
//}
//
//// aesDecryptStream 流式AES解密
//func (c *Crypto) aesDecryptStream() *Crypto {
//	key := c.padKey(c.key, 32)
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		c.Error = fmt.Errorf("failed to create AES cipher: %v", err)
//		return c
//	}
//
//	iv := key[:aes.BlockSize]
//	stream := cipher.NewCFBDecrypter(block, iv)
//
//	buffer := make([]byte, c.chunkSize)
//	for {
//		n, err := c.reader.Read(buffer)
//		if err != nil && err != io.EOF {
//			c.Error = fmt.Errorf("failed to read file: %v", err)
//			return c
//		}
//
//		if n == 0 {
//			break
//		}
//
//		plaintext := make([]byte, n)
//		stream.XORKeyStream(plaintext, buffer[:n])
//
//		if _, err := c.writer.Write(plaintext); err != nil {
//			c.Error = fmt.Errorf("failed to write decrypted data: %v", err)
//			return c
//		}
//	}
//
//	return c
//}

// padKey 填充密钥到指定长度
func (c *Crypto) padKey(key []byte, size int) []byte {
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
func (c *Crypto) Base64Decode() *Crypto {
	if c.Error != nil {
		return c
	}

	if len(c.src) == 0 {
		c.Error = fmt.Errorf("content is empty")
		return c
	}

	decoded, err := base64.StdEncoding.DecodeString(string(c.src))
	if err != nil {
		c.Error = fmt.Errorf("base64 decode failed: %v", err)
		return c
	}

	c.dst = decoded
	return c
}

// HexDecode 十六进制解码
func (c *Crypto) HexDecode() *Crypto {
	if c.Error != nil {
		return c
	}

	if len(c.src) == 0 {
		c.Error = fmt.Errorf("content is empty")
		return c
	}

	decoded, err := hex.DecodeString(string(c.src))
	if err != nil {
		c.Error = fmt.Errorf("hex decode failed: %v", err)
		return c
	}

	c.dst = decoded
	return c
}

//
//// GetFileHash 快速获取文件哈希（便捷方法）
//func GetFileHash(filePath string, algorithm string) (string, error) {
//	Crypto := NewCrypto().FromFile(filePath)
//
//	switch algorithm {
//	case "md5":
//		Crypto.ByMd5()
//	case "sha256":
//		Crypto.BySha256()
//	default:
//		return "", fmt.Errorf("unsupported algorithm: %s", algorithm)
//	}
//
//	if Crypto.Error != nil {
//		return "", Crypto.Error
//	}
//
//	return Crypto.ToHex(), nil
//}

// // FromString encodes from string.
//
//	func (e Crypto) FromString(s string) Crypto {
//		c.src, _ = Var.ToByte(s)
//		return c
//	}
//
// // FromBytes encodes from byte slicc.
//
//	func (e Crypto) FromBytes(b []byte) Crypto {
//		c.src = b
//		return c
//	}
//
// // FromFile encodes from filc.
//
//	func (e Crypto) FromFile(f fs.File) Crypto {
//		c.reader = f
//		return c
//	}
//
// // ToString outputs as string.
//
//	func (e Crypto) ToString() string {
//		res, _ := Var.ToString(c.dst)
//		return res
//	}
//
// // ToBytes outputs as byte slicc.
//
//	func (e Crypto) ToBytes() []byte {
//		if len(c.dst) == 0 {
//			return []byte{}
//		}
//		return c.dst
//	}
//
// //func (e Crypto) ToBase64String() string {
// //	return coding.NewEncoder().FromBytes(c.dst).ByBase64().ToString()
// //}
// //
// //// ToBase64Bytes outputs as base64 byte slicc.
// //func (e Crypto) ToBase64Bytes() []byte {
// //	return coding.NewEncoder().FromBytes(c.dst).ByBase64().ToBytes()
// //}
// //
// //// ToHexString outputs as hex string.
// //func (e Crypto) ToHexString() string {
// //	return coding.NewEncoder().FromBytes(c.dst).ByHex().ToString()
// //}
// //
// //// ToHexBytes outputs as hex byte slicc.
// //func (e Crypto) ToHexBytes() []byte {
// //	return coding.NewEncoder().FromBytes(c.dst).ByHex().ToBytes()
// //}
//
//	func (e Crypto) stream(fn func(io.Writer) io.WriteCloser) ([]byte, error) {
//		var buf bytes.Buffer
//		encoder := fn(&buf)
//
//		// Try to reset the reader position if it's a seeker
//		if seeker, ok := c.reader.(io.Seeker); ok {
//			_, _ = seeker.Seek(0, io.SeekStart)
//		}
//		if _, err := io.CopyBuffer(encoder, c.reader, make([]byte, BufferSize)); err != nil && err != io.EOF {
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
// //	func (e Crypto) ByHex() Crypto {
// //		if c.Error != nil {
// //			return c
// //		}
// //
// //		// Streaming encoding mode
// //		if c.reader != nil {
// //			c.dst, c.Error = c.stream(func(w io.Writer) io.WriteCloser {
// //				return hex.NewStreamEncoder(w)
// //			})
// //			return c
// //		}
// //
// //		// Standard encoding mode
// //		if len(c.src) > 0 {
// //			c.dst = hex.NewStdEncoder().Encode(c.src)
// //		}
// //
// //		return c
// //	}
func (c *Crypto) indirect(in any) any {
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
