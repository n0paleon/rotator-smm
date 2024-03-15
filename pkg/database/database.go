package database

import (
	"RotatorSMM/pkg/database/migrations"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func Init() {
	dsn := "kabarcil_rotator-smm:RotatorSMM3301@tcp(turuhost.my.id:3306)/kabarcil_rotator-smm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Atur level log di sini
	})
	if err != nil {
		panic("failed to connect database")
	}

	db, err := DB.DB()
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(1000)

	migrations.SetDB(DB)
}

func Close() error {
    db, err := DB.DB()
    if err != nil {
        return err
    }
    return db.Close()
}
