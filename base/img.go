package base

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

type imgcircle struct {
	p image.Point
	r int
}

func (c *imgcircle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *imgcircle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *imgcircle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

// 生产放缩后的图片
func ImgReset(src, desc string, size uint) {
	file, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		img, err = jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
	}
	m := resize.Resize(size, size, img, resize.Lanczos3)

	out, err := os.Create(desc)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	png.Encode(out, m)
}

// 生成放缩后的圆形图片
func ImgResetCircle(src, desc string, size uint) {
	file, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		img, err = jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
	}
	sw := size
	sh := size
	if img.Bounds().Dx() > img.Bounds().Max.X {
		sw = size / uint(img.Bounds().Max.Y) * uint(img.Bounds().Max.X)
	} else {
		sh = size / uint(img.Bounds().Max.X) * uint(img.Bounds().Max.Y)
	}

	m := resize.Resize(sw, sh, img, resize.Lanczos3)
	dst := image.NewNRGBA(image.Rect(0, 0, int(size), int(size)))
	draw.DrawMask(dst, dst.Bounds(), m, image.ZP, &imgcircle{image.Pt(int(size/2), int(size/2)), int(size / 2)}, image.ZP, draw.Over)

	out, err := os.Create(desc)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	png.Encode(out, m)
}
