package action

import "github.com/ghf-go/goapp/base"

func IosIconAction() {
	src := base.GetArg(1)
	base.ImgReset(src, "icon.png", 40)
	base.ImgReset(src, "icon@2x.png", 80)
	base.ImgReset(src, "icon@3x.png", 120)

	base.ImgReset(src, "set.png", 29)
	base.ImgReset(src, "set@2x.png", 58)
	base.ImgReset(src, "set@3x.png", 87)

	base.ImgReset(src, "notify.png", 20)
	base.ImgReset(src, "notify@2x.png", 40)
	base.ImgReset(src, "notify@3x.png", 60)
}
