package models

import "github.com/ghf-go/goapp/gtools/app/utils"

type BaseModel interface {
	GetDesc() string
	GetUsage() string
	Run(args *utils.Args) error
}
