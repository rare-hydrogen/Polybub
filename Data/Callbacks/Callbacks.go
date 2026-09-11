package Callbacks

import (
	"Polybub/Data/Audit"
	"Polybub/Utilities/Logger"
	"context"
	"time"
)

type AuditType int

const (
	CREATE AuditType = iota
	UPDATE
	DELETE
)

func AddAuditFields(ctx context.Context, auditType AuditType, af *Audit.AuditFields) {
	// Get data used to fill audit fields
	rdv := ctx.Value("requestDetails").(Logger.RequestDetails)

	// For the given field type, fill the data
	if auditType == CREATE {
		af.CreatedAt = time.Now()
		af.CreatedBy = rdv.UserName
		af.CreatedById = rdv.UserId
	}
	if auditType == UPDATE {
		af.UpdatedAt = time.Now()
		af.UpdatedBy = rdv.UserName
		af.UpdatedById = rdv.UserId
	}
	if auditType == DELETE {
		af.DeletedAt = time.Now()
		af.DeletedBy = rdv.UserName
		af.DeletedById = rdv.UserId
		af.IsDeleted = 1
	}
}

func AddSystemAuditFields(auditType AuditType, af *Audit.AuditFields) {
	// For the given field type, fill the data
	if auditType == CREATE {
		af.CreatedAt = time.Now()
		af.CreatedBy = "System"
		af.CreatedById = 0
	}
	if auditType == UPDATE {
		af.UpdatedAt = time.Now()
		af.UpdatedBy = "System"
		af.UpdatedById = 0
	}
	if auditType == DELETE {
		af.DeletedAt = time.Now()
		af.DeletedBy = "System"
		af.DeletedById = 0
		af.IsDeleted = 1
	}
}
