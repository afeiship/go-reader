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
  "github.com/afeiship/go-reader"
)

func main() {
  // type: "file", "http", "base64"
  r := reader.NewReader("file","README.md")
  content, err := r.Read()
  if err!= nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(content))
}
```