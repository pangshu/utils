package Crypto

import "io"

type Signer struct {
	data   []byte
	sign   []byte
	reader io.Reader
	Error  error
}

type Verifier struct {
	data   []byte
	sign   []byte
	verify bool
	reader io.Reader
	Error  error
}

var BufferSize = 64 * 1024
