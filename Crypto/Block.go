package Crypto

import (
	"crypto/cipher"
	"errors"
	"fmt"
)

// NewCBCEncrypter encrypts data using Cipher Block Chaining (CBC) mode.
// CBC mode encrypts each block of plaintext by XORing it with the previous
// ciphertext block before applying the block cipher algorithm.
func (c *Crypto) NewCBCEncrypter(src, iv []byte, block cipher.Block) (dst []byte, err error) {
	if len(iv) == 0 {
		return dst, errors.New(fmt.Sprintf("iv cannot be empty in '%s' block mode", CBC))
	}

	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return dst, errors.New(fmt.Sprintf("iv length %d must equal block size %d in '%s' block mode", len(iv), blockSize, CBC))
	}

	if len(src)%blockSize != 0 {
		return dst, errors.New(fmt.Sprintf("src length %d must be a multiple of block size %d in '%s' block mode", len(src), blockSize, CBC))
	}

	// Perform CBC encryption using the standard library implementation
	dst = make([]byte, len(src))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(dst, src)
	return
}

// NewCBCDecrypter decrypts data using Cipher Block Chaining (CBC) mode.
// CBC decryption reverses the encryption process by applying the block cipher
// and then XORing with the previous ciphertext block.
func (c *Crypto) NewCBCDecrypter(src, iv []byte, block cipher.Block) (dst []byte, err error) {
	if len(iv) == 0 {
		return dst, errors.New(fmt.Sprintf("iv cannot be empty in '%s' block mode", CBC))
	}

	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return dst, errors.New(fmt.Sprintf("iv length %d must equal block size %d in '%s' block mode", len(iv), blockSize, CBC))
	}

	if len(src)%blockSize != 0 {
		return dst, errors.New(fmt.Sprintf("src length %d must be a multiple of block size %d in '%s' block mode", len(src), blockSize, CBC))
	}

	// Perform CBC decryption using the standard library implementation
	dst = make([]byte, len(src))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(dst, src)
	return
}

// NewCTREncrypter encrypts data using Counter (CTR) mode.
// CTR mode transforms a block cipher into a stream cipher by encrypting
// a counter value and XORing the result with the plaintext.
func (c *Crypto) NewCTREncrypter(src, iv []byte, block cipher.Block) (dst []byte, err error) {
	if len(iv) == 0 {
		return dst, errors.New(fmt.Sprintf("iv cannot be empty in '%s' block mode", CTR))
	}

	// Handle nonce for CTR mode
	// If IV is 12 bytes (nonce), pad it to 16 bytes with zeros
	// This matches Python's pycryptodome behavior
	ctrIV := iv
	if len(iv) == 12 {
		ctrIV = make([]byte, 16)
		copy(ctrIV, iv)
		// The remaining 4 bytes are set to zero (counter starts at 0)
	}

	// Perform CTR encryption using the standard library implementation
	dst = make([]byte, len(src))
	cipher.NewCTR(block, ctrIV).XORKeyStream(dst, src)
	return
}

// NewCTRDecrypter decrypts data using Counter (CTR) mode.
// In CTR mode, decryption is identical to encryption since it's a stream cipher.
func (c *Crypto) NewCTRDecrypter(src, iv []byte, block cipher.Block) (dst []byte, err error) {
	if len(iv) == 0 {
		return dst, errors.New(fmt.Sprintf("iv cannot be empty in '%s' block mode", CTR))
	}

	// Handle nonce for CTR mode
	// If IV is 12 bytes (nonce), pad it to 16 bytes with zeros
	// This matches Python's pycryptodome behavior
	ctrIV := iv
	if len(iv) == 12 {
		ctrIV = make([]byte, 16)
		copy(ctrIV, iv)
		// The remaining 4 bytes are set to zero (counter starts at 0)
	}

	// Perform CTR decryption using the standard library implementation
	dst = make([]byte, len(src))
	cipher.NewCTR(block, ctrIV).XORKeyStream(dst, src)
	return
}

// NewECBEncrypter encrypts data using Electronic Codebook (ECB) mode.
// ECB mode encrypts each block of plaintext independently using the same key.
// Note: ECB mode is generally not recommended for secure applications due to
// its vulnerability to pattern analysis.
func (c *Crypto) NewECBEncrypter(src []byte, block cipher.Block) (dst []byte, err error) {
	blockSize := block.BlockSize()
	if len(src)%blockSize != 0 {
		return dst, errors.New(fmt.Sprintf("src length %d must be a multiple of block size %d in '%s' block mode", len(src), blockSize, ECB))
	}

	// Perform ECB encryption - encrypt each block independently
	dst = make([]byte, len(src))
	for i := 0; i < len(src); i += blockSize {
		block.Encrypt(dst[i:i+blockSize], src[i:i+blockSize])
	}
	return
}

// NewECBDecrypter decrypts data using Electronic Codebook (ECB) mode.
// ECB decryption decrypts each block independently.
func (c *Crypto) NewECBDecrypter(src []byte, block cipher.Block) (dst []byte, err error) {
	blockSize := block.BlockSize()
	if len(src)%blockSize != 0 {
		return dst, errors.New(fmt.Sprintf("src length %d must be a multiple of block size %d in '%s' block mode", len(src), blockSize, ECB))
	}

	// Perform ECB decryption - decrypt each block independently
	dst = make([]byte, len(src))
	for i := 0; i < len(src); i += blockSize {
		block.Decrypt(dst[i:i+blockSize], src[i:i+blockSize])
	}
	return
}

// NewGCMEncrypter encrypts data using Galois/Counter Mode (GCM).
// GCM is an authenticated encryption mode that provides both confidentiality
// and authenticity. It combines CTR mode encryption with a Galois field
// multiplication for authentication.
func (c *Crypto) NewGCMEncrypter(src, nonce, aad []byte, block cipher.Block) (dst []byte, err error) {
	if len(nonce) == 0 {
		return dst, errors.New(fmt.Sprintf("nonce cannot be empty in '%s' block mode", GCM))
	}

	// Create GCM cipher from the underlying block cipher
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return dst, errors.New(fmt.Sprintf("failed to create cipher in '%s' block mode: %v", GCM, err))
	}

	// Perform GCM encryption with authentication
	dst = gcm.Seal(nil, nonce, src, aad)
	return
}

// NewGCMDecrypter decrypts data using Galois/Counter Mode (GCM).
// GCM decryption verifies the authentication tag before decrypting the data.
func (c *Crypto) NewGCMDecrypter(src, nonce, aad []byte, block cipher.Block) (dst []byte, err error) {
	if len(nonce) == 0 {
		return dst, errors.New(fmt.Sprintf("nonce cannot be empty in '%s' block mode", GCM))
	}

	// Create GCM cipher from the underlying block cipher
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return dst, errors.New(fmt.Sprintf("failed to create cipher in '%s' block mode: %v", GCM, err))
	}

	// Perform GCM decryption with authentication verification
	return gcm.Open(nil, nonce, src, aad)
}

// NewCFBEncrypter encrypts data using Cipher Feedback (CFB) mode.
// CFB mode transforms a block cipher into a stream cipher by encrypting
// the previous ciphertext block and XORing the result with the plaintext.
func (c *Crypto) NewCFBEncrypter(src, iv []byte, block cipher.Block) (dst []byte, err error) {
	//if len(iv) == 0 {
	//	return dst, errors.New(fmt.Sprintf("iv cannot be empty in '%s' block mode", CFB))
	//}

	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		iv = c.paddingIVWithZero(iv, blockSize)
	}

	// Perform CFB encryption using the standard library implementation
	dst = make([]byte, len(src))
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(dst, src)
	return
}

func (c *Crypto) paddingIVWithZero(iv []byte, size int) []byte {
	if len(iv) > size {
		//超出，进行截取
		iv = iv[:size]
	} else {
		//不足，进行填充
		if c.IvPaddingType == Left {
			iv = append(make([]byte, size-len(iv)), iv...)
		} else {
			iv = append(iv, make([]byte, size-len(iv))...)
		}
	}

	return iv
}

// NewCFBDecrypter decrypts data using Cipher Feedback (CFB) mode.
// In CFB mode, decryption is identical to encryption since it's a stream cipher.
func (c *Crypto) NewCFBDecrypter(src, iv []byte, block cipher.Block) (dst []byte, err error) {
	if len(iv) == 0 {
		return dst, errors.New(fmt.Sprintf("iv cannot be empty in '%s' block mode", CFB))
	}

	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return dst, errors.New(fmt.Sprintf("iv length %d must equal block size %d in '%s' block mode", len(iv), blockSize, CFB))
	}

	// Perform CFB decryption using the standard library implementation
	dst = make([]byte, len(src))
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(dst, src)
	return
}

// NewOFBEncrypter encrypts data using Output Feedback (OFB) mode.
// OFB mode transforms a block cipher into a stream cipher by repeatedly
// encrypting the initialization vector and using the output as a keystream.
func (c *Crypto) NewOFBEncrypter(src, iv []byte, block cipher.Block) (dst []byte, err error) {
	if len(iv) == 0 {
		return dst, errors.New(fmt.Sprintf("iv cannot be empty in '%s' block mode", OFB))
	}

	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return dst, errors.New(fmt.Sprintf("iv length %d must equal block size %d in '%s' block mode", len(iv), blockSize, OFB))
	}

	// Perform OFB encryption using the standard library implementation
	dst = make([]byte, len(src))
	cipher.NewOFB(block, iv).XORKeyStream(dst, src)
	return
}

// NewOFBDecrypter decrypts data using Output Feedback (OFB) mode.
// In OFB mode, decryption is identical to encryption since it's a stream cipher.
func (c *Crypto) NewOFBDecrypter(src, iv []byte, block cipher.Block) (dst []byte, err error) {
	if len(iv) == 0 {
		return dst, errors.New(fmt.Sprintf("iv cannot be empty in '%s' block mode", OFB))
	}

	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return dst, errors.New(fmt.Sprintf("iv length %d must equal block size %d in '%s' block mode", len(iv), blockSize, OFB))
	}

	// Perform OFB decryption using the standard library implementation
	dst = make([]byte, len(src))
	cipher.NewOFB(block, iv).XORKeyStream(dst, src)
	return
}
