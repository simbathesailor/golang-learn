package cfg

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func IntializeDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "sqlserver://root:sasd@localhost:1433?database=trello_app")
	if err != nil {
		fmt.Println("err is ", err)

	}
	// defer db.Close()
	return db, err
}
