package Crypto

import (
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"strconv"
)

// WithKey 设置加密密钥，支持任意类型
func (e *Cipher) WithKey(value any) *Cipher {
	if value == nil {
		return e
	}

	e.Key = e.anyToBytes(value)
	return e
}

// WithIv 设置偏移量，支持任意类型
func (e *Cipher) WithIv(value any) *Cipher {
	if value == nil {
		return e
	}

	e.Iv = e.anyToBytes(value)
	return e
}

// WithNonce 设置随机数，支持任意类型
func (e *Cipher) WithNonce(value any) *Cipher {
	if value == nil {
		return e
	}

	e.Nonce = e.anyToBytes(value)
	return e
}

// WithPadding 设置附加数据，支持任意类型
func (e *Cipher) WithPadding(value PaddingMode) *Cipher {
	e.Padding = value
	return e
}

// WithBlock 设置附加数据，支持任意类型
func (e *Cipher) WithBlock(value BlockMode) *Cipher {
	e.Block = value
	return e
}

// Encrypt encrypts the source data using the specified cipher.
func (c *Cipher) Encrypt(src []byte, block cipher.Block) (dst []byte, err error) {
	if c.Block == CFB {
		return NewCFBEncrypter(src, c.Iv, block)
	}
	if c.Block == OFB {
		return NewOFBEncrypter(src, c.Iv, block)
	}
	if c.Block == CTR {
		return NewCTREncrypter(src, c.Iv, block)
	}
	if c.Block == GCM {
		return NewGCMEncrypter(src, c.Nonce, c.Addl, block)
	}

	// Only CBC/ECB block modes require add padding
	paddedSrc, err := c.padding(src, block.BlockSize())
	if err != nil {
		return
	}
	if c.Block == CBC {
		return NewCBCEncrypter(paddedSrc, c.Iv, block)
	}
	if c.Block == ECB {
		return NewECBEncrypter(paddedSrc, block)
	}

	return dst, UnsupportedBlockModeError{mode: c.Block}
}

// Decrypt decrypts the source data using the specified cipher.
func (c *Cipher) Decrypt(src []byte, block cipher.Block) (dst []byte, err error) {
	if c.Block == CFB {
		return NewCFBDecrypter(src, c.Iv, block)
	}
	if c.Block == OFB {
		return NewOFBDecrypter(src, c.Iv, block)
	}
	if c.Block == CTR {
		return NewCTRDecrypter(src, c.Iv, block)
	}
	if c.Block == GCM {
		return NewGCMDecrypter(src, c.Nonce, c.Addl, block)
	}

	// Only CBC/ECB block modes require remove padding
	var paddedDst []byte
	if c.Block == CBC {
		if paddedDst, err = NewCBCDecrypter(src, c.Iv, block); err != nil {
			return
		}
		return c.unpadding(paddedDst)
	}
	if c.Block == ECB {
		if paddedDst, err = NewECBDecrypter(src, block); err != nil {
			return
		}
		return c.unpadding(paddedDst)
	}

	return dst, UnsupportedBlockModeError{mode: c.Block}
}

// padding adds padding to the source data.
func (c *Cipher) padding(src []byte, blockSize int) (dst []byte, err error) {
	switch c.Padding {
	case No:
		return NewNoPadding(src), nil
	case Zero:
		return NewZeroPadding(src, blockSize), nil
	case PKCS5:
		return NewPKCS5Padding(src), nil
	case PKCS7:
		return NewPKCS7Padding(src, blockSize), nil
	case AnsiX923:
		return NewAnsiX923Padding(src, blockSize), nil
	case ISO97971:
		return NewISO97971Padding(src, blockSize), nil
	case ISO10126:
		return NewISO10126Padding(src, blockSize), nil
	case ISO78164:
		return NewISO78164Padding(src, blockSize), nil
	case Bit:
		return NewBitPadding(src, blockSize), nil
	default:
		return dst, UnsupportedPaddingModeError{mode: c.Padding}
	}
}

// unpadding removes padding from the source data.
func (c *Cipher) unpadding(src []byte) (dst []byte, err error) {
	switch c.Padding {
	case No:
		return NewNoUnPadding(src), nil
	case Zero:
		return NewZeroUnPadding(src), nil
	case PKCS5:
		return NewPKCS5UnPadding(src), nil
	case PKCS7:
		return NewPKCS7UnPadding(src), nil
	case AnsiX923:
		return NewAnsiX923UnPadding(src), nil
	case ISO97971:
		return NewISO97971UnPadding(src), nil
	case ISO10126:
		return NewISO10126UnPadding(src), nil
	case ISO78164:
		return NewISO78164UnPadding(src), nil
	case Bit:
		return NewBitUnPadding(src), nil
	default:
		return dst, UnsupportedPaddingModeError{mode: c.Padding}
	}
}

// anyToBytes 将任意类型转换为字节数组
func (e *Cipher) anyToBytes(value any) []byte {
	if value == nil {
		return nil
	}

	// 根据不同类型转换为字节数组
	switch v := value.(type) {
	case string:
		return []byte(v)
	case []byte:
		return v
	case int, int8, int16, int32, int64:
		return []byte(fmt.Sprintf("%d", v))
	case uint, uint8, uint16, uint32, uint64:
		return []byte(fmt.Sprintf("%d", v))
	case float32, float64:
		return []byte(fmt.Sprintf("%f", v))
	case bool:
		return []byte(strconv.FormatBool(v))
	default:
		// 对于其他类型，尝试JSON序列化
		jsonData, err := json.Marshal(v)
		if err != nil {
			return nil
		}
		return jsonData
	}

}
