package models

import "time"

// 工程类型
const (
	ProjectTypeMobile  = "mobile"  // 移动应用
	ProjectTypeDesktop = "desktop" // 桌面应用
	ProjectTypeLinux   = "linux"   // linux应用
	ProjectTypeWeb     = "web"     // web应用
)

// ProjectModel 工程表，创建/修改时间由数据库维护
type ProjectModel struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"size:128;not null" json:"name"`
	GitUrl      string    `gorm:"size:512;not null" json:"gitUrl"`
	Type        string    `gorm:"size:32;not null;default:web" json:"type"`
	TestBranch  string    `gorm:"size:128" json:"testBranch"`
	TestVersion string    `gorm:"size:32" json:"testVersion"`
	ProdBranch  string    `gorm:"size:128" json:"prodBranch"`
	ProdVersion string    `gorm:"size:32" json:"prodVersion"`
	CreatedAt   time.Time `gorm:"autoCreateTime:false;default:CURRENT_TIMESTAMP(3)" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"<-:false;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updatedAt"`
}

// TableName 表名
func (ProjectModel) TableName() string {
	return "project"
}
