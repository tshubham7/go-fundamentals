package base64_decoder

import (
	"encoding/base64"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// save file to given path
func ToFile(b64, file_path string) error {
	// remove the extra content
	b64 = strings.Split(b64, "base64,")[1]

	dec, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return err
	}

	mimeType := http.DetectContentType(dec)

	var file_name string
	switch mimeType {
	case "image/png":
		file_name = "file.png"
	case "image/jpeg":
		file_name = "file.jpeg"
	case "application/pdf":
		file_name = "file.pdf"
	}
	f, err := os.Create(filepath.Join(file_path, file_name))

	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return err
	}

	// Sync commits the current contents of the file to stable storage.
	// Typically, this means flushing the file system's in-memory copy of recently written data to disk.
	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}
