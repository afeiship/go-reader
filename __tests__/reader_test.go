package reader

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/afeiship/go-reader"
)

func TestLocalReader(f *testing.T) {
	localFileReader, err := reader.NewFileReader("file", "hello.txt")
	if err != nil {
		fmt.Println("Error reading local file:", err)
	} else {
		io.Copy(os.Stdout, localFileReader)
	}
}

func TestHttpReader(f *testing.T) {
	httpFileReader, err := reader.NewFileReader("http", "https://web-assets.alo7.com/assets/text/hello.txt")
	if err != nil {
		fmt.Println("Error reading http file:", err)
	} else {
		io.Copy(os.Stdout, httpFileReader)
	}
}

// b64Str := "SGVsbG8gd29ybGQh" // "Hello world!" 的 Base64 编码
func TestBase64Reader(f *testing.T) {
	b64Str := "SGVsbG8gd29ybGQh"
	base64FileReader, err := reader.NewFileReader("base64", b64Str)
	if err != nil {
		fmt.Println("Error reading base64 file:", err)
	} else {
		io.Copy(os.Stdout, base64FileReader)
	}
}
