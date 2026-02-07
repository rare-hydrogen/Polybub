INSERT INTO Users (FirstName,LastName,Username,Password,Salt,AccountEmail,AccountPhone,UserGroup,TotpKey,CreatedAt,UpdatedAt,DeletedAt,CreatedBy,UpdatedBy,DeletedBy) VALUES
	 ('Admin','User','admin123','','','email-here','phone-here',1,'','2025-12-03 03:59:05.79891864+00:00','2025-12-22 06:18:08.993918875+00:00',NULL,'System','System',NULL);

INSERT INTO Permissions (UserId,Name,IsCreate,IsRead,IsUpdate,IsDelete,CreatedAt,UpdatedAt,DeletedAt,CreatedBy,UpdatedBy,DeletedBy) VALUES
	 (1,'Dashboard',1,1,1,1,NULL,NULL,NULL,NULL,NULL,NULL),
	 (1,'MfaCode',1,1,1,1,NULL,NULL,NULL,NULL,NULL,NULL),
	 (1,'Users',1,1,1,1,NULL,NULL,NULL,NULL,NULL,NULL),
