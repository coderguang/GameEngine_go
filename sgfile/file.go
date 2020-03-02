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

func GetPath(path string) (string, error) {
	fliterlist := strings.Split(path, "/")
	if len(fliterlist) > 0 {
		pathlist := fliterlist[0 : len(fliterlist)-1]
		str := ""
		for _, v := range pathlist {
			str += v + "/"
		}
		return str, nil
	}
	return "", errors.New("error path,path=" + path)
}

func WriteFile(path string, filename string, datas []byte) (int, string, error) {
	AutoMkDir(path)
	finalFileName := path + filename
	newFile, err := Create(finalFileName)
	if err != nil {
		return 0, finalFileName, err
	}
	defer newFile.Close()
	writeNum, err := newFile.Write(datas)
	if err != nil || newFile.Close() != nil {
		return writeNum, finalFileName, err
	}
	return writeNum, finalFileName, nil
}

func GetAllFile(pathname string) ([]string, error) {
	fileList := []string{}
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			dirlist, _ := GetAllFile(pathname + fi.Name() + "/")
			for _, v := range dirlist {
				fileList = append(fileList, v)
			}
		} else {
			fileList = append(fileList, pathname+fi.Name())
		}
	}
	return fileList, err
}
