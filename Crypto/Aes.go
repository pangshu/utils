package Crypto

import (
	"crypto/aes"
)

// ByAes encrypts by aes.
func (e *Encrypter) ByAes(c *Cipher) *Encrypter {
	if e.Error != nil {
		return e
	}

	if len(e.src) > 0 {
		block, err := aes.NewCipher(c.Key)
		if err != nil {
			e.Error = err
			return e
		}
		e.dst, e.Error = c.Encrypt(e.src, block)
	}
	return e
}

//
//// ByAes decrypts by aes.
//func (d Decrypter) ByAes(c *cipher.AesCipher) Decrypter {
//	if d.Error != nil {
//		return d
//	}
//
//	// Streaming decryption mode
//	if d.reader != nil {
//		d.dst, d.Error = d.stream(func(r io.Reader) io.Reader {
//			return aes.NewStreamDecrypter(r, c)
//		})
//		return d
//	}
//
//	// Standard decryption mode
//	if len(d.src) > 0 {
//		d.dst, d.Error = aes.NewStdDecrypter(c).Decrypt(d.src)
//	}
//
//	return d
//}
