package learn_gorm

import (
	"testing"
)
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email *string
	Age   uint8
}

func TestConnect(t *testing.T) {
	dsn := "root:zhaoxin@tcp(121.199.12.107:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	sqlOpen := mysql.Open(dsn)
	db, _ := gorm.Open(sqlOpen, &gorm.Config{})
	db.AutoMigrate(&User{})

	//for i := 0; i < 10; i++ {
	//	db.Save(&User{})
	//}
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})

}
