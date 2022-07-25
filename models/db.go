package models

/*
	数据库连接实现，包括基础的数据库连接以及uuid生成和哈希加密
*/

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/jabin/Chatplatm/config"
)

var Db *sql.DB

func init() {
	var err error

	config := LoadConfig()
	driver := config.Db.Driver
	source := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true", config.Db.User, config.Db.Password, config.Db.Address, config.Db.Database)
	Db, err = sql.Open(driver, source)
	if err != nil {
		log.Fatal(err)
	}
}

// 通过RFC 创建UUID
func createUUID() string {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)

	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return uuid

}

// 哈希加密
func Encrypt(plaintext string) string {
	crypttext := fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return crypttext
}
