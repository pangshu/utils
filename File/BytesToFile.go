package File

import (
	"os"
)

// BytesToFile 将字节切片写入文件
func BytesToFile(data []byte, filePath string) error {
	// 创建或截断文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将数据写入文件
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}