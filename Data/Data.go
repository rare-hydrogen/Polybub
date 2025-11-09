package Data

import (
	"Polybub/Utilities"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	gormConfig := &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	db, err := gorm.Open(sqlite.Open(Utilities.GlobalConfig.Connection), gormConfig)
	if err != nil {
		panic("failed to connect to database")
	}

	db.Exec("PRAGMA foreign_keys = ON")

	return db
}
