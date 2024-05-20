package global

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"goprojects/pkg/zap/config"
)

var (
	Server      config.Server
	LOG         *zap.Logger
	JwtSecret   string
	DB          *gorm.DB
)