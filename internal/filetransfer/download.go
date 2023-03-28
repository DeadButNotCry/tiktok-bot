package filetransfer

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func FileTransferIoDownload(link string) {

	// TODO: Parse link and get id for download
	id := link
	f, _ := os.Create(id)
	r, _ := http.Get("https://filetransfer.io/data-package/o4JdqBkB/download")
	_, err := io.Copy(f, r.Body)
	if err != nil {
		fmt.Println(err)
	}
}
