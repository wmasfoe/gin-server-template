package global

import (
	"go-blog-server/config"
	"gorm.io/gorm"
)

var (
	CONFIG *config.Config
	DB     *gorm.DB
)
