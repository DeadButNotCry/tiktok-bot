package filetransfer

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func UploadToAnonfiles(path string) string {
	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)
	fw, err := writer.CreateFormFile("file", filepath.Base(path))
	check(err)

	fd, err := os.Open(path)
	check(err)

	defer fd.Close()
	_, err = io.Copy(fw, fd)
	check(err)

	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.anonfiles.com/upload", form)
	check(err)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	check(err)

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	check(err)

	if resp.StatusCode == 200 {
		var respStruct AnonFilesResponse
		json.Unmarshal(bodyText, &respStruct)
		log.Println(respStruct.Data.File.URL.Full)
		return respStruct.Data.File.URL.Full
	} else {
		return "error"
	}
}
