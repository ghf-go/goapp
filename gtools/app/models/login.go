package models

import (
	"errors"

	"github.com/ghf-go/goapp/gtools/app/utils"
)

type LoginModel struct {
}

func NewLoginModel() *LoginModel {
	return &LoginModel{}
}

func (a *LoginModel) GetDesc() string {
	return "登陆"
}
func (a *LoginModel) GetUsage() string {
	return `-h 服务器地址 -t token   #登陆 `
}
func (a *LoginModel) Run(args *utils.Args) error {
	host := args.Get("h")
	token := args.Get("t")
	if host == "" || token == "" {
		return errors.New("host or token is empty")
	}

	return utils.SaveConfig(&utils.Config{
		Host:  host,
		Token: token,
	})
}
