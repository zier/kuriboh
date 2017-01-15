package main

import (
	"github.com/zier/kuriboh/config"
	"github.com/zier/kuriboh/downloader"
	"github.com/zier/kuriboh/niceoppai"
	"github.com/zier/kuriboh/service"
)

func main() {
	sc := niceoppai.New()
	dl := downloader.New()
	c := config.New()

	k := service.New(c, sc, dl)
	k.Start()
}
