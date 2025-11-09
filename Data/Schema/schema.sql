PRAGMA foreign_keys = ON;

CREATE TABLE FooBar (
    Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    Name TEXT NOT NULL,
    Type TEXT NOT NULL,
    CreatedAt DATETIME,
    UpdatedAt DATETIME,
    DeletedAt DATETIME,
    CreatedBy DATETIME,
    UpdatedBy DATETIME,
    DeletedBy DATETIME
)