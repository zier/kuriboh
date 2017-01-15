package service

// NiceoppaiScrap ...
type NiceoppaiScrap interface {
	GetImagesPathFromCartoonName(cartoonName string, chapter int) ([]string, error)
}

// Downloader ...
type Downloader interface {
	Image(url, pathName, fileName string) error
}
