package converter

import (
	"bytes"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"io"
)

// Big5ToUtf8 将Big5编码的字符串转换为UTF-8编码的字符串 / converts Big5 encoded string to UTF-8 encoded
func (c *Converter) Big5ToUtf8(v []byte) ([]byte, error) {
	// 创建一个转换读取器，将Big5编码转换为UTF-8编码
	// bytes.NewReader(v) 创建一个从v读取的Reader
	// traditionalchinese.Big5.NewDecoder() 创建一个Big5解码器
	reader := transform.NewReader(bytes.NewReader(v), traditionalchinese.Big5.NewDecoder())
	// 读取转换后的所有数据并返回W
	return io.ReadAll(reader)
}
