package action

import "github.com/ghf-go/goapp/base"

func WxAppImg() {
	src := base.GetArg(1)
	base.ImgReset(src, "wechatapp28.png", 28)
	base.ImgReset(src, "wechatapp108.png", 108)
}
