package main

import (
	"github.com/zier/kuriboh/config"
	"github.com/zier/kuriboh/niceoppai"
	"github.com/zier/kuriboh/service"
)

func main() {
	scrapper := niceoppai.New()
	c := config.New()
	k := service.New(c, scrapper)
	k.Start()
}
