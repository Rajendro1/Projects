package includes

import (
	"database/sql"
	"log"
)

var err error

func Connect() {
	db, err = sql.Open("mysql", DB_USERNAME+":"+DB_PASSWORD+"@tcp("+DB_HOST+":3306)/"+DB_NAME)
	if err != nil {
		log.Println("Error To Connect Databae")
	}
}
