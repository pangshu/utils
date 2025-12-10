package Crypto

var BufferSize = 64 * 1024

// Crypto 密码结构体
type Crypto struct {
	src   []byte //原始内容
	dst   []byte //目标内容
	Error error

	Key           []byte      //密钥
	Iv            []byte      //初始向量
	IvPaddingType PaddingType //初始向量填充类型，0向后填充，1向前填充
	Additional    []byte      //附加数据
	Nonce         []byte      //随机数
	Padding       PaddingMode //填充数据
	Block         BlockMode
}

//type Cipher struct {
//	Key           []byte      //密钥
//	Iv            []byte      //初始向量
//	IvPaddingType PaddingType //初始向量填充类型
//	Addl          []byte      //附加数据
//	Nonce         []byte      //随机数
//	Padding       PaddingMode //填充数据
//	Block         BlockMode
//	Error         error
//}

type PaddingType string

const (
	Left  PaddingType = "Left"  // Left padding - fills with zeros on left
	Right PaddingType = "Right" // Right padding - fills with zeros on right
)

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

type Option func(*Crypto)

func Init(opts ...Option) *Crypto {
	cfg := &Crypto{
		IvPaddingType: Right,
		Padding:       No,
		Block:         CBC,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

// WithKey 设置加密密钥
func WithKey(v any) Option {
	return func(c *Crypto) {
		c.Key = c.anyToBytes(v)
	}
}

// WithIv 添加偏移量
func WithIv(v any) Option {
	return func(c *Crypto) {
		c.Iv = c.anyToBytes(v)
	}
}

// WithBlock 添加块模式
func WithBlock(v BlockMode) Option {
	return func(c *Crypto) {
		c.Block = v
	}
}

// WithNonce 添加随机数
func WithNonce(v any) Option {
	return func(c *Crypto) {
		c.Nonce = c.anyToBytes(v)
	}
}

// WithPadding 添加填充数据
func WithPadding(v PaddingMode) Option {
	return func(c *Crypto) {
		c.Padding = v
	}
}

func WithIvPadding(v PaddingType) Option {
	return func(c *Crypto) {
		c.IvPaddingType = v
	}
}
