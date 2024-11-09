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

// Options 结构体封装了读取文件所需的配置信息
type Options struct {
	Type   string
	Source any
}

// NewReader 创建 FileReader，根据传入的类型动态选择读取方式
func NewReader(opts *Options) (*FileReader, error) {
	var reader io.Reader
	var err error

	if !isValideFileType(opts.Type) {
		return nil, errors.New("unsupported file type")
	}

	switch opts.Type {
	case "reader":
		reader = opts.Source.(io.Reader)
	case "file":
		reader, err = createLocalFileReader(opts.Source.(string))
	case "http":
		reader, err = createHttpFileReader(opts.Source.(string))
	case "base64":
		reader, err = createB64FileReader(opts.Source.(string))
	case "bytes":
		reader = createBytesReader(opts.Source.([]byte))
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

func createBytesReader(data []byte) io.Reader {
	return bytes.NewReader(data)
}

func isValideFileType(fileType string) bool {
	switch fileType {
	case "file", "http", "base64", "bytes", "reader":
		return true
	default:
		return false
	}
}

// Read 实现 io.Reader 接口方法，读取 FileReader 中的数据
func (f *FileReader) Read(p []byte) (n int, err error) {
	return f.reader.Read(p)
}
