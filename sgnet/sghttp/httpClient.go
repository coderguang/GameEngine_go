package sghttp

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgstring"
)

//formName is for server to get file by r.FormFile(formName)
func PostFile(filename string, target_url string, formName string) (*http.Response, error) {

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// use the body_writer to write the Part headers to the buffer

	_, err := body_writer.CreateFormFile(formName, filename)
	if err != nil {
		sglog.Error("post File error,file:", filename, ",url:", target_url, err)
		return nil, err
	}

	fh, err := os.Open(filename)
	if err != nil {
		sglog.Error("open file failed!", filename, err)
		return nil, err
	}
	// need to know the boundary to properly close the part myself.
	boundary := body_writer.Boundary()

	//close_string := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

	// use multi-reader to defer the reading of the file data until
	// writing to the socket buffer.
	request_reader := io.MultiReader(body_buf, fh, close_buf)
	fi, err := fh.Stat()
	if err != nil {
		fmt.Printf("Error Stating file: %s", filename)
		return nil, err
	}
	req, err := http.NewRequest("POST", target_url, request_reader)
	if err != nil {
		return nil, err
	}

	// Set headers for multipart, and Content Length
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = fi.Size() + int64(body_buf.Len()) + int64(close_buf.Len())

	return http.DefaultClient.Do(req)

}

func PostMultiFormFile(url string, formMap map[string]string, values map[string]string) (*http.Response, error) {

	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)
	randomStr := sgstring.RandStringAndNumRunes(15)
	body_writer.SetBoundary(randomStr)

	//field

	if values != nil {
		for k, v := range values {
			body_writer.WriteField(k, v)
		}
	}

	//form data
	if formMap != nil {
		for k, v := range formMap {
			_, err := body_writer.CreateFormFile(k, v)
			if err != nil {
				sglog.Error("PostMultiFormFile create form file error,key:", k, ",file:", v, err)
				return nil, err
			}
			fileBytes, err := ioutil.ReadFile(v)
			if err != nil {
				sglog.Error("PostMultiFormFile read file error key:", k, ",file:", v, err)
				return nil, err
			}
			_, err = body_buf.Write(fileBytes)
			if err != nil {
				sglog.Error("PostMultiFormFile write file error key:", k, ",file:", v, err)
				return nil, err
			}
		}
	}

	body_writer.Close()

	req_reader := io.MultiReader(body_buf)
	req, err := http.NewRequest("POST", url, req_reader)
	if err != nil {
		sglog.Error("PostMultiFormFile create request error", err)
		return nil, err
	}
	// 添加Post头
	// req.Header.Set("Connection", "close")
	// req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Content-Type", body_writer.FormDataContentType())
	req.ContentLength = int64(body_buf.Len())

	return http.DefaultClient.Do(req)
}
