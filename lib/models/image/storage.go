package image

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path"

	"github.com/nfnt/resize"

	"github.com/astaxie/beego"

	"github.com/ikeikeikeike/gopkg/convert"
	encimg "github.com/ikeikeikeike/gopkg/encoding/image"

	"bitbucket.org/ikeikeikeike/antenna/models"
)

func SaveImage(m *models.Image, r io.ReadSeeker, mime string, filename string) (err error) {
	var (
		ext string
		img image.Image
	)
	if ext, err = encimg.ImageExt(filename, mime); err != nil {
		return
	}
	if img, err = encimg.Decord(r, ext); err != nil {
		return
	}

	m.Name = path.Base(filename)
	m.Width = img.Bounds().Dx()
	m.Height = img.Bounds().Dy()

	if err = m.Insert(); err != nil || m.Id <= 0 {
		return
	}

	path := GenImagePath(m)
	os.MkdirAll(path, 0755)

	fullPath := GenImageFilePath(m, 0)
	if _, err = r.Seek(0, 0); err != nil {
		return
	}

	var (
		f    *os.File
		file *os.File
	)
	if f, err = os.OpenFile(fullPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644); err != nil {
		return
	} else {
		file = f
	}

	if _, err = io.Copy(file, r); err != nil {
		os.RemoveAll(fullPath)
		return
	}

	if ext != ".gif" {
		if m.Width > models.ImageSizeSmall {
			if err = ImageResize(m, img, models.ImageSizeSmall); err != nil {
				os.RemoveAll(fullPath)
				return
			}
		}
		if m.Width > models.ImageSizeMiddle {
			if err = ImageResize(m, img, models.ImageSizeMiddle); err != nil {
				os.RemoveAll(fullPath)
				return
			}
		}
	}
	return
}

func ImageResize(img *models.Image, im image.Image, width int) error {
	savePath := GenImageFilePath(img, width)
	im = resize.Resize(uint(width), 0, im, resize.Bilinear)

	var file *os.File
	if f, err := os.OpenFile(savePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644); err != nil {
		return err
	} else {
		file = f
	}
	defer file.Close()

	var err error
	switch img.Ext {
	case ".jpg":
		err = jpeg.Encode(file, im, &jpeg.Options{90})
	case ".png":
		err = png.Encode(file, im)
	default:
		err = fmt.Errorf("<ImageResize> unsupport image format")
	}

	return err
}

func GenImagePath(img *models.Image) string {
	return "upload/img/" + beego.Date(img.Created, "y/m/d/s/") + convert.ToStr(img.Id) + "/"
}

func GenImageFilePath(img *models.Image, width int) string {
	var size string
	if width == 0 {
		size = "full"
	} else {
		size = convert.ToStr(width)
	}
	return GenImagePath(img) + size + img.Ext
}
