package Crypto

var BufferSize = 64 * 1024

// Encrypter 加密器结构体
type Encrypter struct {
	src   []byte //原始内容
	dst   []byte //目标内容
	Error error
}

type Cipher struct {
	Key     []byte      //密钥
	Iv      []byte      //初始向量
	Addl    []byte      //附加数据
	Nonce   []byte      //随机数
	Padding PaddingMode //填充数据
	Block   BlockMode
}
type PaddingMode string

const (
	No       PaddingMode = "No"        // No padding - data must be exact block size
	Zero     PaddingMode = "Zero"      // Zero padding - fills with zeros, always adds padding
	PKCS5    PaddingMode = "PKCS5"     // PKCS5 padding - RFC 2898, 8-byte blocks only
	PKCS7    PaddingMode = "PKCS7"     // PKCS7 padding - RFC 5652, variable block size
	AnsiX923 PaddingMode = "AnsiX.923" // ANSI X.923 padding - zeros + length byte
	ISO97971 PaddingMode = "ISO9797-1" // ISO/IEC 9797-1 padding method 1
	ISO10126 PaddingMode = "ISO10126"  // ISO/IEC 10126 padding - random + length byte
	ISO78164 PaddingMode = "ISO7816-4" // ISO/IEC 7816-4 padding - same as ISO9797-1
	Bit      PaddingMode = "Bit"       // Bit padding - 0x80 + zeros
)

type BlockMode string

// Supported block cipher modes
const (
	CBC BlockMode = "CBC" // Cipher Block Chaining mode
	ECB BlockMode = "ECB" // Electronic Codebook mode
	CTR BlockMode = "CTR" // Counter mode
	GCM BlockMode = "GCM" // Galois/Counter Mode
	CFB BlockMode = "CFB" // Cipher Feedback mode
	OFB BlockMode = "OFB" // Output Feedback mode
)
