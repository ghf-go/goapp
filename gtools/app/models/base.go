package models

type BaseModel interface {
	GetDesc() string
	GetUsage() string
	Run() error
}
