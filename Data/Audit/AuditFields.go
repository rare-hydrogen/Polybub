package Audit

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type AuditFields struct {
	CreatedAt   time.Time             `gorm:"column:created_at;" json:"CreatedAt"`
	UpdatedAt   time.Time             `gorm:"column:updated_at;" json:"UpdatedAt"`
	DeletedAt   time.Time             `gorm:"column:deleted_at;" json:"DeletedAt"`
	CreatedBy   string                `gorm:"column:created_by;" json:"CreatedBy"`
	UpdatedBy   string                `gorm:"column:updated_by;" json:"UpdatedBy"`
	DeletedBy   string                `gorm:"column:deleted_by;" json:"DeletedBy"`
	CreatedById int32                 `gorm:"column:created_by_id;" json:"CreatedById"`
	UpdatedById int32                 `gorm:"column:updated_by_id;" json:"UpdatedById"`
	DeletedById int32                 `gorm:"column:deleted_by_id;" json:"DeletedById"`
	IsDeleted   soft_delete.DeletedAt `gorm:"column:is_deleted;softDelete:flag" json:"IsDeleted"`
}
