package service

import (
	"fmt"

	errors "github.com/kataras/go-errors"
	"github.com/zier/kuriboh/config"
)

// Kuriboh ...
type Kuriboh struct {
	*config.Config
	Scrapper   NiceoppaiScrap
	Downloader Downloader
}

// LoadResponse ...
type LoadResponse struct {
	Message string
	Err     error
}

// New ...
func New(c *config.Config, ns NiceoppaiScrap, dl Downloader) *Kuriboh {
	return &Kuriboh{c, ns, dl}
}

// ValidateParams ...
func (k *Kuriboh) ValidateParams() error {
	if k.CartoonProfile.DestinationPath == "" {
		return errors.New("DestinationPath is requried")
	}

	if k.CartoonProfile.CartoonPath == "" {
		return errors.New("CartoonPath is requried")
	}

	if k.CartoonProfile.StartChapter == 0 {
		return errors.New("StartChapter is requried")
	}

	if k.CartoonProfile.EndChapter == 0 {
		return errors.New("EndChapter is requried")
	}

	return nil
}

// Start ...
func (k *Kuriboh) Start() error {
	if err := k.ReadCLIParams(); err != nil {
		return err
	}

	if err := k.ValidateParams(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(fmt.Sprintf("Start download %s chapter %d - %d", k.CartoonProfile.CartoonPath, k.CartoonProfile.StartChapter, k.CartoonProfile.EndChapter))

	chapterCount := k.Config.CartoonProfile.EndChapter - k.Config.CartoonProfile.StartChapter

	cs := make(chan LoadResponse)

	for i := 0; i <= chapterCount; i++ {
		go k.Load(cs, k.CartoonProfile.CartoonPath, k.Config.CartoonProfile.StartChapter+i)
	}

	for i := 0; i <= chapterCount; i++ {
		lr := <-cs
		fmt.Println(lr.Message)
	}

	return nil
}

// Load ...
func (k *Kuriboh) Load(cs chan LoadResponse, cartoonName string, chapterNumber int) {
	listImages, err := k.Scrapper.GetImagesPathFromCartoonName(cartoonName, chapterNumber)
	if err != nil {
		cs <- LoadResponse{fmt.Sprintf("Error load images url from chapter %d", chapterNumber), err}
	}

	for pageNumber, urlPath := range listImages {
		if err := k.Downloader.Image(urlPath, fmt.Sprintf("./%s/%d/", cartoonName, chapterNumber), fmt.Sprintf("%d.jpg", pageNumber)); err != nil {
			cs <- LoadResponse{fmt.Sprintf("Error load page %d", pageNumber), err}
		}
	}

	cs <- LoadResponse{fmt.Sprintf("Chapter %d load succeed!", chapterNumber), err}
}
