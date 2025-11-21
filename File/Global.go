package File

import (
	"io"
	"os"
)

// ReadFile 读取文件内容并返回字节切片（别名）
func ReadFile(filePath string) ([]byte, error) {
	return FileToBytes(filePath)
}

// WriteFile 将字节切片写入文件（别名）
func WriteFile(data []byte, filePath string) error {
	return BytesToFile(data, filePath)
}

// FileToBytesByBuffer 使用缓冲方式读取大文件并返回字节切片
func FileToBytesByBuffer(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 使用 CopyN 限制读取大小，防止内存溢出
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ReaderToBytes 从 io.Reader 读取数据并返回字节切片
func ReaderToBytes(reader io.Reader) ([]byte, error) {
	return io.ReadAll(reader)
}

// BytesToWriter 将字节切片写入 io.Writer
func BytesToWriter(data []byte, writer io.Writer) (int, error) {
	return writer.Write(data)
}