package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

func NewPostgresDB(cfg DatabaseConfig) (*gorm.DB, error) {
	dsn := cfg.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移模型
	// db.AutoMigrate(&models.User{})

	return db, nil
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func (c DatabaseConfig) GetDSN() string {
	return "host=" + c.Host + " user=" + c.User + " password=" + c.Password +
		" dbname=" + c.Name + " port=" + strconv.Itoa(c.Port) + " sslmode=disable TimeZone=Asia/Shanghai"
}
