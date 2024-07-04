package base

import (
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

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
func ImgResetCircle(src, desc string, size int32) {

}
