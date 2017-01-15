package service

import (
	"fmt"

	"github.com/zier/kuriboh/config"
)

// Kuriboh ...
type Kuriboh struct {
	*config.Config
	Scrapper   NiceoppaiScrap
	Downloader Downloader
}

// New ...
func New(c *config.Config, ns NiceoppaiScrap, dl Downloader) *Kuriboh {
	return &Kuriboh{c, ns, dl}
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
		if err := k.Downloader.Image(urlPath, fmt.Sprintf("./%s/%d/", cartoonName, chapterNumber), fmt.Sprintf("%d.jpg", pageNumber)); err != nil {
			return err
		}

	}

	return nil
}
