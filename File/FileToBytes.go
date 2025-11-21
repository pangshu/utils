package File

import (
	"io"
	"os"
)

// FileToBytes 读取文件内容并返回字节切片
func FileToBytes(filePath string) ([]byte, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 读取文件所有内容
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}