package data

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"os"
	"encoding/json"
	"fmt"
	"crypto/sha1"
	"crypto/rand"
	"casual-talk/utils"
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
	// Must to be append '?parseTime=true' after the 'dataSourceName',
	// or it will occur 'unsupported Scan, storing driver.Value type []uint8 into type *time.Time'
	dataSourceName := fmt.Sprintf("%s:%s@/%s?parseTime=true", info.Username, info.Password, info.Database)
	Db, err = sql.Open(info.DriverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return
}


func loadSqlInfo() {
	file, err := os.Open("data/sql.json")
	if err != nil {
		utils.Danger("Cannot open sql file", err)
	}
	decoder := json.NewDecoder(file)
	info = sqlInfo{}
	err = decoder.Decode(&info)
	if err != nil {
		utils.Danger("Cannot get configuration from file", err)
	}
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		utils.Danger("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1
func Encrypt(plainText string) (encryptedText string) {
	encryptedText = fmt.Sprintf("%x", sha1.Sum([]byte(plainText)))
	return
}
