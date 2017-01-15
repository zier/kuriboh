package niceoppai

import (
	"net/http"

	"fmt"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Scrapper ...
type Scrapper struct {
	sourceURL string
}

// New ...
func New() *Scrapper {
	return &Scrapper{
		sourceURL: "http://www.niceoppai.net/",
	}
}

// GetImagesPathFromCartoonName ...
func (s *Scrapper) GetImagesPathFromCartoonName(cartoonName string, chapter int) ([]string, error) {
	images := []string{}

	resp, err := http.Get(fmt.Sprintf("%s%s/%d/?all", s.sourceURL, cartoonName, chapter))
	if err != nil {
		return nil, err
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.Img && n.Parent.Parent.Attr[0].Val == "wpm_pag mng_rdr" && n.Attr[0].Val != "img_mng_enl" {
			return true
		}
		return false
	}

	imgs := scrape.FindAll(root, matcher)
	for _, img := range imgs {
		images = append(images, scrape.Attr(img, "src"))
	}

	return images, nil
}
