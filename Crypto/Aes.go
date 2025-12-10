package Crypto

import (
	"crypto/aes"
)

// ByAes encrypts by aes.
func (c *Crypto) ByAes() *Crypto {
	if c.Error != nil {
		return c
	}

	if len(c.src) == 0 {
		return c
	}

	if len(c.src) > 0 {
		block, err := aes.NewCipher(c.Key)
		if err != nil {
			c.Error = err
			return c
		}
		//fmt.Println("++++++++++++++")
		//fmt.Println(block.BlockSize())
		//iv := c.Key[:block.BlockSize()]
		//fmt.Println(string(iv))
		//fmt.Println("++++++++++++++")
		c.dst, c.Error = c.Encrypt(c.src, block)
	}
	return c
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
