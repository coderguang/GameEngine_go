package sghttp

import (
	"errors"
	"io/ioutil"
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

func CheckIsAllowFiles(w http.ResponseWriter, r *http.Request, filelist []string) (multipart.File, string, error) {
	for _, v := range filelist {
		file, _, err := r.FormFile(v)
		if err == nil {
			return file, v, nil
		}
	}
	return nil, "", errors.New("INVALID_FILE_UPLOAD no support upload file type")
}

func CheckFileTypeMatch(w http.ResponseWriter, r *http.Request, fileType string, file multipart.File) ([]byte, error) {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return fileBytes, err
	}
	detectedFileType := http.DetectContentType(fileBytes)
	if detectedFileType != fileType {
		return fileBytes, errors.New("INVALID_FILE_TYPE not match with" + fileType)
	}
	return fileBytes, nil
}
