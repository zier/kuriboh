package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Downloader ...
type Downloader struct {
}

// New ...
func New() *Downloader {
	return &Downloader{}
}

// Image ...
func (d *Downloader) Image(url, pathName, fileName string) error {
	os.MkdirAll(pathName, os.ModePerm)

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	file, err := os.Create(fmt.Sprintf("%s%s", pathName, fileName))
	if err != nil {
		return err
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	file.Close()

	return nil
}
