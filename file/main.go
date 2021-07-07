package main

import (
	"github.com/tshubham7/go-fundamentals/file/base64_decoder"
	"github.com/tshubham7/go-fundamentals/file/file_encoder"
)

func main() {
	// getting base64
	base64, err := file_encoder.ToBase64("/Users/shubhamdhanera/Downloads/file.pdf")
	if err != nil {
		panic(err)
	}
	// saving base64 as a file
	err = base64_decoder.ToFile(base64, "./files")
	if err != nil {
		panic(err)
	}
}
