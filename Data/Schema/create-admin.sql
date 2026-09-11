INSERT INTO users (first_name,last_name,username,password,salt,account_email,account_phone,user_group,totp_key,created_at,updated_at,deleted_at,created_by,updated_by,deleted_by,created_by_id,updated_by_id,deleted_by_id,is_deleted) VALUES
	 ('Admin','User','admin123','','','email-here','phone-here',1,'','0001-01-01 00:00:00+00:00','0001-01-01 00:00:00+00:00','0001-01-01 00:00:00+00:00','System','','',1,0,0,0);

INSERT INTO permissions (user_id,name,is_create,is_read,is_update,is_delete,created_at,updated_at,deleted_at,created_by,updated_by,deleted_by,created_by_id,updated_by_id,deleted_by_id,is_deleted) VALUES
	 (1,'Dashboard',true,true,true,true,'0001-01-01 00:00:00+00:00','0001-01-01 00:00:00+00:00','0001-01-01 00:00:00+00:00','System','','',1,0,0,0),
	 (1,'MfaCode',true,true,true,true,'0001-01-01 00:00:00+00:00','0001-01-01 00:00:00+00:00','0001-01-01 00:00:00+00:00','System','','',1,0,0,0),
	 (1,'Users',true,true,true,true,'0001-01-01 00:00:00+00:00','0001-01-01 00:00:00+00:00','0001-01-01 00:00:00+00:00','System','','',1,0,0,0);
