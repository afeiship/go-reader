# go-reader
> File reader for go.

## installation
```sh
go get -u github.com/afeiship/go-reader
```

## usage
```go
package main

import (
	"fmt"
	"io"
	"github.com/afeiship/go-reader"
)

func main() {
    // type: "file", "http", "base64"
    fileReader, _ := reader.NewReader("file", "file.txt")
    data, _ := io.ReadAll(fileReader)
    fmt.Println(string(data))
}
```