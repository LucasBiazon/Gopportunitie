package handler

import (
	"github.com/LucasBiazon/Gopportunitie.git/config"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *config.Logger
)

func InitializeHandler() {
	db = config.GetDataBase()
	logger = config.GetLogger("handler")
}
