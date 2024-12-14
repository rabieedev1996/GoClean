package GoClean_Common

import (
	"fmt"
	"mime"
	"path"
	"path/filepath"
	"strings"
)

type Path string

func (r *Path) GetExtension() *string {
	parts := strings.Split(fmt.Sprint(*r), ".")
	if len(parts) == 0 {
		return nil
	} else {
		return &parts[1]
	}
}

func (r *Path) GetContentType() *string {
	parts := strings.Split(fmt.Sprint(*r), ".")
	var ext string
	if len(parts) == 0 {
		return nil
	} else if len(parts) == 1 {
		ext = parts[0]
	} else {
		ext = parts[1]

	}
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	contentType := mime.TypeByExtension(ext)
	return &contentType
}

func (r *Path) FileName() *string {
	rawPath := string(*r)

	if strings.Contains(rawPath, "://") {
		base := path.Base(rawPath)
		return &base
	}

	base := filepath.Base(rawPath)
	return &base
}
