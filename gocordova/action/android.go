package action

import "github.com/ghf-go/goapp/base"

func AndroidIconAction() {
	src := base.GetArg(1)
	base.ImgReset(src, "ic_launcher_l.png", 36)
	base.ImgReset(src, "ic_launcher_m.png", 48)
	base.ImgReset(src, "ic_launcher_h.png", 72)
	base.ImgReset(src, "ic_launcher_xh.png", 96)
	base.ImgReset(src, "ic_launcher_xxh.png", 144)
	base.ImgReset(src, "ic_launcher_xxxh.png", 192)
}

// cordova 生成 icon 图标
func AndroidCordovaIconAction() {
	src := base.GetArg(1)
	base.ImgReset(src, "platforms/android/app/src/main/res/mipmap-ldpi/ic_launcher.png", 36)
	base.ImgReset(src, "platforms/android/app/src/main/res/mipmap-mdpi/ic_launcher.png", 36)
	base.ImgReset(src, "platforms/android/app/src/main/res/mipmap-hdpi/ic_launcher.png", 48)
	base.ImgReset(src, "platforms/android/app/src/main/res/mipmap-xhdpi/ic_launcher.png", 72)
	base.ImgReset(src, "platforms/android/app/src/main/res/mipmap-xxhdpi/ic_launcher.png", 96)
	base.ImgReset(src, "platforms/android/app/src/main/res/mipmap-xxxhdpi/ic_launcher.png", 144)
}
