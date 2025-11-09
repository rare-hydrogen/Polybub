package Audit

import (
	"time"

	"gorm.io/gorm"
)

type AuditFields struct {
	CreatedAt *time.Time      `gorm:"<-:created;column:CreatedAt;type:DATETIME" json:"CreatedAt"`
	UpdatedAt *time.Time      `gorm:"<-:updated;column:UpdatedAt;type:DATETIME" json:"UpdatedAt"`
	DeletedAt *gorm.DeletedAt `gorm:"<-:deleted;column:DeletedAt;type:DATETIME" json:"DeletedAt"`
	CreatedBy *string         `gorm:"<-:created;column:CreatedBy;type:DATETIME" json:"CreatedBy"`
	UpdatedBy *string         `gorm:"<-:updated;column:UpdatedBy;type:DATETIME" json:"UpdatedBy"`
	DeletedBy *string         `gorm:"<-:deleted;column:DeletedBy;type:DATETIME" json:"DeletedBy"`
}
