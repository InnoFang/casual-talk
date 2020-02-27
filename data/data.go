package data

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"os"
	"encoding/json"
	"fmt"
)

type sqlInfo struct {
	driverName string
	username string
	password string
	database string
}

var info sqlInfo
var Db *sql.DB

func init() {
	loadConfig()
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@/%s", info.username, info.password, info.database)
	Db, err = sql.Open(info.driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return
}


func loadConfig() {
	file, err := os.Open("sql.json")
	if err != nil {
		log.Fatalln("Cannot open mysql file", err)
	}
	decoder := json.NewDecoder(file)
	info = sqlInfo{}
	err = decoder.Decode(&info)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}
