package action

import (
	"runtime"

	"github.com/ghf-go/goapp/base"
)

// https://developer.android.google.cn/tools/sdkmanager?hl=zh-cn
// https://googledownloads.cn/android/repository/commandlinetools-mac-11076708_latest.zip
// android 环境安装
func androidAction() {
	switch runtime.GOOS {
	case "darwin":
		// base.ShRun("wget https://googledownloads.cn/android/repository/commandlinetools-mac-11076708_latest.zip ; unzip commandlinetools-mac-11076708_latest.zip;mkdir -p ~/sdk;mv cmdline-tools ~/sdk/;rm -rf commandlinetools-mac-11076708_latest.zip;echo 'PATH=$PATH:~/sdk/cmdline-tools/bin' >> ~/.bashrc;echo 'PATH=$PATH:~/sdk/cmdline-tools/bin' >> ~/.zshrc")
		base.ShRun("brew install java gradle ;brew casks install android-commandlinetools  android-sdk android-messages android-ndk  android-platform-tools  androidtool  android-studio flutter")
	case "linux":
	case "win":
	}
}
