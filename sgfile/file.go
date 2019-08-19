package sgfile

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MkdirAll(dir string, mode os.FileMode) error {
	err := os.MkdirAll(dir, mode)
	return err
}

func Create(filename string) (*os.File, error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return f, err
}

func GetFileContentAsStringLines(filePath string) ([]string, error) {
	result := []string{}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("read file:", filePath, " error:", err)
		return result, err
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		result = append(result, lineStr)
	}
	return result, nil
}

func AutoMkDir(path string) {
	if ok, _ := PathExists(path); !ok {
		MkdirAll(path, os.ModePerm)
	}
}

func Rename(oldPath string, newPath string) error {
	err := os.Rename(oldPath, newPath)
	return err
}

func GetFileName(path string) (string, error) {
	fliterlist := strings.Split(path, "/")
	if len(fliterlist) > 0 {
		return fliterlist[len(fliterlist)-1], nil
	}
	return "", errors.New("error path,path=" + path)
}

func GetFileRawName(path string) (string, error) {
	filename, err := GetFileName(path)
	if err != nil {
		return "", err
	}
	fliterlist := strings.Split(path, ".")
	if len(fliterlist) > 0 {

		str := ""
		for k, v := range fliterlist {
			if k == len(fliterlist)-1 {
				continue
			}
			str += v
		}
		return str, nil
	}
	return "", errors.New("error filename,filename=" + filename)
}
