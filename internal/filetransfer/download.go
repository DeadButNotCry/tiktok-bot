package filetransfer

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// https://filetransfer.io/data-package/r1wDDgFv#link
func FileTransferIoDownload(link string) string {
	parts := strings.Split(link, "/")
	id := strings.Split(parts[len(parts)-1], "#")[0]
	f, _ := os.Create(filepath.Join("./videos/", id+".mp4"))
	r, _ := http.Get("https://filetransfer.io/data-package/" + id + "/download")
	_, err := io.Copy(f, r.Body)
	if err != nil {
		log.Println(err)
	}
	return id + ".mp4"
}
