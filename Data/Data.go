package Data

import (
	"Polybub/Utilities"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	if Utilities.GlobalConfig.Env == Utilities.TEST_ENV {
		return getSqliteConnection()
	}

	return getPgConnection()
}

func getPgConnection() *gorm.DB {
	gormConfig := &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	db, err := gorm.Open(postgres.Open(Utilities.GlobalConfig.Connection), gormConfig)
	if err != nil {
		panic(err.Error())
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func getSqliteConnection() *gorm.DB {
	gormConfig := &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	db, err := gorm.Open(sqlite.Open(Utilities.GlobalConfig.Connection), gormConfig)
	if err != nil {
		panic(err.Error())
	}

	db.Exec(`pragma journal_mode = wal;
		pragma synchronous = normal;
		pragma temp_store = memory;
		pragma busy_timeout = 500;	
		pragma foreign_keys = on;`)

	return db
}
