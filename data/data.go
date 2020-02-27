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
	DriverName string
	Username string
	Password string
	Database string
}

var info sqlInfo
var Db *sql.DB

func init() {
	loadSqlInfo()
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@/%s", info.Username, info.Password, info.Database)
	Db, err = sql.Open(info.DriverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return
}


func loadSqlInfo() {
	file, err := os.Open("data/sql.json")
	if err != nil {
		log.Fatalln("Cannot open sql file", err)
	}
	decoder := json.NewDecoder(file)
	info = sqlInfo{}
	err = decoder.Decode(&info)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}
