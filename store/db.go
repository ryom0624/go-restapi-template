package store

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"webapp/config"
)

type DB struct {
	*gorm.DB
}

func InitDB(cfg *config.Config) (*DB, error) {
	var dsn string
	if cfg.GoEnv == "localhost" {
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FTokyo",
			cfg.DBUser,
			cfg.DBPass,
			cfg.DBHost,
			cfg.DBName,
		)
	} else {
		dsn = fmt.Sprintf(
			"%s:%s@unix(%s/%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FTokyo",
			cfg.DBUser,
			cfg.DBPass,
			"/cloudsql",
			cfg.DBHost,
			cfg.DBName,
		)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile), // io writer
			logger.Config{
				SlowThreshold: 200 * time.Millisecond, // Slow SQL threshold
				LogLevel:      logger.Info,            // Log level
				Colorful:      true,                   // Disable color
			},
		),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
