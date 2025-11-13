PRAGMA foreign_keys = ON;

CREATE TABLE Users (
	Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FirstName TEXT NOT NULL,
	LastName TEXT NOT NULL,
    Username TEXT NOT NULL UNIQUE,
	Password TEXT NOT NULL,
	Salt TEXT NOT NULL,
	AccountEmail TEXT NOT NULL,
	AccountPhone TEXT NOT NULL,
	UserGroup INTEGER NOT NULL,
    CreatedAt DATETIME,
    UpdatedAt DATETIME,
    DeletedAt DATETIME,
    CreatedBy TEXT,
    UpdatedBy TEXT,
    DeletedBy TEXT
);

CREATE TABLE UserPasswordResets (
	Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	UserId INTEGER NOT NULL,
	ResetKey TEXT NOT NULL,
    CreatedAt DATETIME,
    UpdatedAt DATETIME,
    DeletedAt DATETIME,
	CreatedBy TEXT,
	UpdatedBy TEXT,
	DeletedBy TEXT,
	CONSTRAINT UserPasswordResets_Users_FK FOREIGN KEY (UserId) REFERENCES Users(Id)
);

CREATE TABLE Permissions (
	Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	UserId INTEGER NOT NULL,
	Name TEXT NOT NULL,
	IsCreate INTEGER NOT NULL,
	IsRead INTEGER NOT NULL,
	IsUpdate INTEGER NOT NULL,
	IsDelete INTEGER NOT NULL,
    CreatedAt DATETIME,
    UpdatedAt DATETIME,
    DeletedAt DATETIME,
	CreatedBy TEXT,
	UpdatedBy TEXT,
	DeletedBy TEXT,
	CONSTRAINT Permissions_User_FK FOREIGN KEY (UserId) REFERENCES Users(Id)
);

CREATE TABLE FooBar (
    Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    Name TEXT NOT NULL,
    Type TEXT NOT NULL,
    CreatedAt DATETIME,
    UpdatedAt DATETIME,
    DeletedAt DATETIME,
    CreatedBy TEXT,
    UpdatedBy TEXT,
    DeletedBy TEXT
)