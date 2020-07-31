package cfg

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func IntializeDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:Stack@123@/trello_app?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("err is ", err)

	} else {
		fmt.Println("Database connection successfull")
	}
	// v := db.Ping()
	// if v != nil {
	// 	panic(v.Error()) // proper error handling instead of panic in your app
	// }

	// defer db.Close()
	return db, err
}
