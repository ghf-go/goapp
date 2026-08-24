package models

import "time"

// 构建记录状态
const (
	BuildStatusPending = "pending"
	BuildStatusRunning = "running"
	BuildStatusSuccess = "success"
	BuildStatusFailed  = "failed"
)

// 构建环境
const (
	BuildEnvTest = "test"
	BuildEnvProd = "prod"
)

// BuildRecordModel 构建记录表，创建/修改时间由数据库维护
type BuildRecordModel struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ProjectID uint64    `gorm:"not null;index" json:"projectId"`
	Env       string    `gorm:"size:16;not null" json:"env"`
	Branch    string    `gorm:"size:128" json:"branch"`
	Version   string    `gorm:"size:32" json:"version"`
	Status    string    `gorm:"size:16;not null;default:pending" json:"status"`
	Log       string    `gorm:"type:longtext" json:"log"`
	CreatedAt time.Time `gorm:"autoCreateTime:false;default:CURRENT_TIMESTAMP(3)" json:"createdAt"`
	UpdatedAt time.Time `gorm:"<-:false;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updatedAt"`
}

// TableName 表名
func (BuildRecordModel) TableName() string {
	return "build_record"
}
