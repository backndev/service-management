package config

import (
	"backend-onboarding/model/entity"
	_ "github.com/apache/calcite-avatica-go/v5"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func ConnectMySQL() (*gorm.DB, error) {
	connectionString := CONFIG["MYSQL_USER"] + ":" + CONFIG["MYSQL_PASS"] + "@tcp(" + CONFIG["MYSQL_HOST"] + ":" + CONFIG["MYSQL_PORT"] + ")/" + CONFIG["MYSQL_SCHEMA"] + "?parseTime=true"
	mysqlConn, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Println("Error connect to MySQL: ", err.Error())
		return nil, err
	}
	log.Println("MySQL connection success")

	sqlDB, errDB := mysqlConn.DB()
	if errDB != nil {
		log.Println(errDB)
	} else {
		sqlDB.SetMaxIdleConns(2)
		sqlDB.SetMaxOpenConns(1000)
	}

	mysqlConn.AutoMigrate(&entity.User{})
	mysqlConn.AutoMigrate(&entity.Role{})
	mysqlConn.AutoMigrate(&entity.Product{})

	role := entity.Role{}
	roleList := mysqlConn.Model(&entity.Role{}).Select("ID").Scan(&role)
	if roleList.RowsAffected == 0 {
		var role = [5]string{"admin", "maker", "checker", "signer", "viewer"}
		for i := 0; i < len(role); i++ {
			ID := uuid.New()
			mysqlConn.Create(&entity.Role{ID: ID, Title: role[i], Active: false})
		}
	}

	return mysqlConn, nil
}
