package reader

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"os"
)

// FileReader 结构体封装了读取文件数据的能力
type FileReader struct {
	reader io.Reader
}

// ReaderOptions 结构体封装了读取文件所需的配置信息
type ReaderOptions struct {
	fileType string
	source   string
}

// NewReader 创建 FileReader，根据传入的类型动态选择读取方式
func NewReader(opts *ReaderOptions) (*FileReader, error) {
	var reader io.Reader
	var err error

	if !isValideFileType(opts.fileType) {
		return nil, errors.New("unsupported file type")
	}

	switch opts.fileType {
	case "file":
		reader, err = createLocalFileReader(opts.source)
	case "http":
		reader, err = createHttpFileReader(opts.source)
	case "base64":
		reader, err = createB64FileReader(opts.source)
	default:
		err = errors.New("unsupported file type")
	}

	if err != nil {
		return nil, err
	}
	return &FileReader{reader: reader}, nil
}

// createLocalFileReader 创建一个本地文件 reader
func createLocalFileReader(filename string) (io.Reader, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

// createHttpFileReader 创建一个 HTTP 文件 reader
func createHttpFileReader(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

// createB64FileReader 创建一个 Base64 字符串 reader
func createB64FileReader(b64 string) (io.Reader, error) {
	data, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

func isValideFileType(fileType string) bool {
	switch fileType {
	case "file", "http", "base64":
		return true
	default:
		return false
	}
}

// Read 实现 io.Reader 接口方法，读取 FileReader 中的数据
func (f *FileReader) Read(p []byte) (n int, err error) {
	return f.reader.Read(p)
}
