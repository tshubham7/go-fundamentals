package file_encoder

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

// read file from the path and encode it to base64
func ToBase64(file_path string) (string, error) {
	bytes, err := ioutil.ReadFile(file_path)
	if err != nil {
		return "", err
	}

	var base64Encoding string

	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/png":
		base64Encoding = "data:file/png;base64,"
	case "image/jpeg":
		base64Encoding = "data:file/jpeg;base64,"
	case "application/pdf":
		base64Encoding = "data:file/pdf;base64,"
	}

	base64Encoding += toBase64(bytes)
	return base64Encoding, err
}
func toBase64(bytes []byte) string {
	// convert bytes into encoded string
	return base64.StdEncoding.EncodeToString(bytes)
}
