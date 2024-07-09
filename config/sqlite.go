package config

import (
	"os"

	"github.com/LucasBiazon/Gopportunitie.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {

	dbPatch := "./db/main.db"
	logger := GetLogger("sqlite")
	_, err := os.Stat(dbPatch)

	if os.IsNotExist(err) {
		logger.Infof("SQLite file not found, creating a new one")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPatch)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPatch), &gorm.Config{})
	if logger != nil {
		logger.Errorf("SQLite opening error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("SQLite migration error: %v", err)
		return nil, err
	}
	return db, nil
}
