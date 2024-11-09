package reader

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/afeiship/go-reader"
)

func TestLocalReader(f *testing.T) {
	opts := reader.Options{
		Type:   "file",
		Source: "hello.txt",
	}
	localFileReader, err := reader.NewReader(&opts)
	if err != nil {
		fmt.Println("Error reading local file:", err)
	} else {
		io.Copy(os.Stdout, localFileReader)
	}
}

func TestHttpReader(f *testing.T) {
	opts := reader.Options{
		Type:   "http",
		Source: "https://web-assets.alo7.com/assets/text/hello.txt",
	}
	httpFileReader, err := reader.NewReader(&opts)
	if err != nil {
		fmt.Println("Error reading http file:", err)
	} else {
		io.Copy(os.Stdout, httpFileReader)
	}
}

// b64Str := "SGVsbG8gd29ybGQh" // "Hello world!" 的 Base64 编码
func TestBase64Reader(f *testing.T) {
	// b64Str := "SGVsbG8gd29ybGQh"
	// base64FileReader, err := reader.NewReader("base64", b64Str)
	opts := reader.Options{
		Type:   "base64",
		Source: "SGVsbG8gd29ybGQh",
	}
	base64FileReader, err := reader.NewReader(&opts)
	if err != nil {
		fmt.Println("Error reading base64 file:", err)
	} else {
		io.Copy(os.Stdout, base64FileReader)
	}
}
