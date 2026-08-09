package models

type LoginModel struct {
}

func NewLoginModel() *LoginModel {
	return &LoginModel{}
}

func (a *LoginModel) GetDesc() string {
	return "登陆"
}
func (a *LoginModel) GetUsage() string {
	return `-t token   #登陆 `
}
func (a *LoginModel) Run() error {
	return nil
}
