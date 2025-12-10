package Crypto

import (
	"crypto/cipher"
	"errors"
	"fmt"
)

// Encrypt encrypts the source data using the specified cipher.
func (c *Crypto) Encrypt(src []byte, block cipher.Block) (dst []byte, err error) {
	if c.Block == CFB {
		return c.NewCFBEncrypter(src, c.Iv, block)
	}
	if c.Block == OFB {
		return c.NewOFBEncrypter(src, c.Iv, block)
	}
	if c.Block == CTR {
		return c.NewCTREncrypter(src, c.Iv, block)
	}
	if c.Block == GCM {
		return c.NewGCMEncrypter(src, c.Nonce, c.Additional, block)
	}

	// Only CBC/ECB block modes require add padding
	paddedSrc, err := c.padding(src, block.BlockSize())
	if err != nil {
		return
	}
	if c.Block == CBC {
		return c.NewCBCEncrypter(paddedSrc, c.Iv, block)
	}
	if c.Block == ECB {
		return c.NewECBEncrypter(paddedSrc, block)
	}

	return dst, errors.New(fmt.Sprintf("unsupported block mode '%s'", c.Block))
}

// Decrypt decrypts the source data using the specified cipher.
func (c *Crypto) Decrypt(src []byte, block cipher.Block) (dst []byte, err error) {
	if c.Block == CFB {
		return c.NewCFBDecrypter(src, c.Iv, block)
	}
	if c.Block == OFB {
		return c.NewOFBDecrypter(src, c.Iv, block)
	}
	if c.Block == CTR {
		return c.NewCTRDecrypter(src, c.Iv, block)
	}
	if c.Block == GCM {
		return c.NewGCMDecrypter(src, c.Nonce, c.Additional, block)
	}

	// Only CBC/ECB block modes require remove padding
	var paddedDst []byte
	if c.Block == CBC {
		if paddedDst, err = c.NewCBCDecrypter(src, c.Iv, block); err != nil {
			return
		}
		return c.unpadding(paddedDst)
	}
	if c.Block == ECB {
		if paddedDst, err = c.NewECBDecrypter(src, block); err != nil {
			return
		}
		return c.unpadding(paddedDst)
	}

	return dst, errors.New(fmt.Sprintf("unsupported block mode '%s'", c.Block))
}
