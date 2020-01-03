package sghttp

import (
	"errors"
	"mime/multipart"
	"net/http"
)

//maxSize KB
func CheckFileMaxSize(w http.ResponseWriter, r *http.Request, maxSize int64) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxSize*1024)
	if err := r.ParseMultipartForm(maxSize); err != nil {
		return errors.New("FILE_TO_BIG")
	}
	return nil
}

func GetFileDetectContentType(file multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
