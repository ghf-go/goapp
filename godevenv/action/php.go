package action

import (
	"runtime"

	"github.com/ghf-go/goapp/base"
)

func phpAction() {
	switch runtime.GOOS {
	case "darwin":
		base.ShRun("brew install nginx httpd apache-httpd apache2  redis mysql@8.0 brew-php-switcher shivammathur/php/php@7.1 shivammathur/php/php@7.4 shivammathur/php/php@8.0 shivammathur/php/php@8.4 phpunit composer;brew casks install sequel-ace")
	case "linux":
	case "win":
	}
}
