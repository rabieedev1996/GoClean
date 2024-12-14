package GoClean_Common

import (
	"io"
	"math/rand"
	"mime/multipart"
	"strings"
	"time"
)

var contentTypeToExtension = map[string]string{
	"image/jpeg":               ".jpg",
	"image/png":                ".png",
	"text/plain":               ".txt",
	"application/json":         ".json",
	"application/pdf":          ".pdf",
	"image/gif":                ".gif",
	"audio/mpeg":               ".mp3",
	"video/mp4":                ".mp4",
	"application/zip":          ".zip",
	"application/msword":       ".doc",
	"application/vnd.ms-excel": ".xls",
}

func RandomNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}

type ContentType string

func (r *ContentType) ToExtension() string {
	if extension, found := contentTypeToExtension[strings.ToLower(string(*r))]; found {
		return extension
	}
	return ""
}

type RestFile multipart.FileHeader

func (r *RestFile) GetDetails() (string, []byte) {
	fileName := r.Filename
	file := (*multipart.FileHeader)(r)
	openedFile, _ := file.Open()

	defer openedFile.Close()

	// خواندن محتویات فایل به صورت آرایه بایت
	fileBytes, _ := io.ReadAll(openedFile)

	return fileName, fileBytes
}
