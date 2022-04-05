package config

import (
	"backend-onboarding/model/entity"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func ConnectPostgre() (*gorm.DB, *sql.DB, error) {

	dsn := fmt.Sprintf(
		`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta`,
		CONFIG["POSTGRES_HOST"],
		CONFIG["POSTGRES_USER"],
		CONFIG["POSTGRES_PASS"],
		CONFIG["POSTGRES_SCHEMA"],
		CONFIG["POSTGRES_PORT"],
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	postgreConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Println("Error connect to Postgre: ", err.Error())
		return nil, nil, err
	}

	sqlDb, errDb := postgreConn.DB()
	if errDb != nil {
		log.Println(errDb)
	} else {
		sqlDb.SetMaxIdleConns(2)
		sqlDb.SetMaxOpenConns(1000)
	}

	log.Println("Postgres connection success")

	postgreConn.AutoMigrate(&entity.User{})
	postgreConn.AutoMigrate(&entity.Role{})
	postgreConn.AutoMigrate(&entity.Product{})

	role := entity.Role{}
	roleList := postgreConn.Model(&entity.Role{}).Select("ID").Scan(&role)
	if roleList.RowsAffected == 0 {
		var role = [5]string{"admin", "maker", "checker", "signer", "viewer"}
		for i := 0; i < len(role); i++ {
			ID := uuid.New()
			postgreConn.Create(&entity.Role{ID: ID, Title: role[i], Active: false})
		}
	}

	return postgreConn, sqlDb, nil
}
