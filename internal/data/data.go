package data

import (
	"eas_api/internal/conf"
	"eas_api/internal/pkg/igorm"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAdministratorRepo)

type Data struct {
	db *gorm.DB
}

func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	// 初始化 GORM
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		Logger: igorm.NewLogger(logger, gormlogger.LogLevel(3)),
	})
	if err != nil {
		return nil, nil, err
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Data{
			db: db,
		}, func() {
			if err := sqlDB.Close(); err != nil {
				log.NewHelper(logger).Error("failed to close database", err)
			}
		}, nil
}

func GetSingleRecordByScope[T any](db *gorm.DB) (*T, error) {
	var result T
	err := db.First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}
