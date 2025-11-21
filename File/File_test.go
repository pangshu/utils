package File

import (
	"bytes"
	"os"
	"testing"
)

func TestFileToBytes(t *testing.T) {
	// 创建一个临时文件用于测试
	content := []byte("Hello, World!")
	tmpfile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// 测试 FileToBytes 函数
	data, err := FileToBytes(tmpfile.Name())
	if err != nil {
		t.Fatalf("FileToBytes failed: %v", err)
	}

	if string(data) != string(content) {
		t.Fatalf("Expected %s, got %s", content, data)
	}
}

func TestBytesToFile(t *testing.T) {
	// 创建测试数据
	content := []byte("Hello, File!")
	tmpfile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// 测试 BytesToFile 函数
	err = BytesToFile(content, tmpfile.Name())
	if err != nil {
		t.Fatalf("BytesToFile failed: %v", err)
	}

	// 验证文件内容
	data, err := os.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != string(content) {
		t.Fatalf("Expected %s, got %s", content, data)
	}
}

func TestReaderToBytes(t *testing.T) {
	// 创建测试数据
	content := []byte("Hello, Reader!")
	reader := bytes.NewReader(content)

	// 测试 ReaderToBytes 函数
	data, err := ReaderToBytes(reader)
	if err != nil {
		t.Fatalf("ReaderToBytes failed: %v", err)
	}

	if string(data) != string(content) {
		t.Fatalf("Expected %s, got %s", content, data)
	}
}

func TestBytesToWriter(t *testing.T) {
	// 创建测试数据
	content := []byte("Hello, Writer!")
	var buffer bytes.Buffer

	// 测试 BytesToWriter 函数
	_, err := BytesToWriter(content, &buffer)
	if err != nil {
		t.Fatalf("BytesToWriter failed: %v", err)
	}

	if buffer.String() != string(content) {
		t.Fatalf("Expected %s, got %s", content, buffer.String())
	}
}