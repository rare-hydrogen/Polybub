package Data

import (
	"Polybub/Auth/OAuth2"
	"Polybub/Data/Callbacks"
	"Polybub/Utilities"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	claims := OAuth2.GlobalClaims
	Callbacks.GlobalByName = claims.Name

	gormConfig := &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	db, err := gorm.Open(sqlite.Open(Utilities.GlobalConfig.Connection), gormConfig)
	if err != nil {
		panic("failed to connect to database")
	}

	Callbacks.SetCallbacks(db)

	db.Exec(`pragma journal_mode = wal;
		pragma synchronous = normal;
		pragma temp_store = memory;
		pragma busy_timeout = 500;	
		pragma foreign_keys = on;`)

	return db
}
