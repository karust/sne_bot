package api

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// UploadImage ... Uploads image to blog using API
func UploadImage(file *os.File, path string) error {
	/*file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	*/
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("imgs", path)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://10.1.1.178:8080/api/image/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "JWT eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjVjMzk5ZTg2ZGY2MDY0MDAyNmM2NzIwNyIsImRpc3BsYXlOYW1lIjoiWW9iYUJvYmEiLCJlbWFpbCI6InFxQHFxLnFxIiwiaWF0IjoxNTQ3MjgwMDExLCJleHAiOjE2MzM2ODAwMTEsImlzcyI6IkF3ZXNvbWVTZXJ2aWNlIn0._9qVvf-efPHjcitA20m_xkNYvgMd8gxKAmTevQi8eyU")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)
		fmt.Println(body)
	}
	return nil
}

func main() {
	/*
		request, err := newfileUploadRequest("http://10.1.1.178:8080/api/image/upload", "imgs", "/tmp/aaaaa.jpg")
		if err != nil {
			log.Fatal(err)
		}
	*/

	file, err := os.Open("/tmp/aaaaa.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	UploadImage(file, "/tmp/aaaaa.jpg")
}
