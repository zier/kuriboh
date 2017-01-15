package service

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/zier/kuriboh/config"
)

// Kuriboh ...
type Kuriboh struct {
	*config.Config
	Scrapper NiceoppaiScrap
}

// New ...
func New(c *config.Config, ns NiceoppaiScrap) *Kuriboh {
	return &Kuriboh{c, ns}
}

// Start ...
func (k *Kuriboh) Start() error {
	if err := k.ReadCLIParams(); err != nil {
		return err
	}

	cartoonName := "toaru_ossan_no_vrmmo_katsudouki"
	chapterNumber := 2

	listImages, err := k.Scrapper.GetImagesPathFromCartoonName(cartoonName, chapterNumber)
	if err != nil {
		return err
	}

	for pageNumber, urlPath := range listImages {
		if err := downloadImage(urlPath, fmt.Sprintf("./%s/%d/", cartoonName, chapterNumber), fmt.Sprintf("%d.jpg", pageNumber)); err != nil {
			return err
		}

	}

	return nil
}

func downloadImage(url, pathName, fileName string) error {
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
