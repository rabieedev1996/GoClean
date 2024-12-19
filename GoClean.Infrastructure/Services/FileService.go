package Services

import (
	GoClean_Common "GoClean/GoClean.Common"
	GoClean_Domain "GoClean/GoClean.Domain"
	"bufio"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type FileService struct {
	byteArray          []byte
	fileType           string
	contentType        string
	loadedStream       bool
	fileManagerBaseUrl string
}

func NewFileService(config GoClean_Domain.Configs_FileService) *FileService {
	return &FileService{}
}

func (r *FileService) LoadFileFromStorage(path string) {
	file, _ := os.Open(path)
	fileStatus, _ := file.Stat()
	r.byteArray = make([]byte, fileStatus.Size())
	bufio.NewReader(file).Read(r.byteArray)
	r.loadedStream = true
	r.contentType = http.DetectContentType(r.byteArray)
	r.fileType = filepath.Ext(path)
}
func (r *FileService) LoadFileFromStream(stream io.Reader, fileType *string, contentType *string) {
	if fileType == nil {
		contentType := GoClean_Common.ContentType(*contentType)
		r.fileType = contentType.ToExtension()
	}
	if contentType == nil {
		path := GoClean_Common.Path(*fileType)
		r.contentType = *path.GetExtension()
	}
	r.byteArray, _ = io.ReadAll(stream)
	r.loadedStream = true
}
func (r *FileService) LoadFileFromByteArray(bytes []byte, fileType *string, contentType *string) {
	if fileType == nil {
		contentType := GoClean_Common.ContentType(*contentType)
		r.fileType = contentType.ToExtension()
	}
	if contentType == nil {
		path := GoClean_Common.Path(*fileType)
		r.contentType = *path.GetContentType()
	}
	r.byteArray = bytes
	r.loadedStream = true
}
