package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" //sql
	uuid "github.com/satori/go.uuid"
)

func db() (database *sql.DB) {
	database, err := sql.Open("sqlite3", "./data/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	return
}

func createUUID() (Uuid string) {
	var e error 
	return uuid.Must(uuid.NewV4(), e).String()
	
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
// func createUUID() (uuid string) {
//   u := new([16]byte)
//   _, err := rand.Read(u[:])
//   if err != nil {
//     log.Fatalln("Cannot generate UUID", err)
//   }

//   // 0x40 is reserved variant from RFC 4122
//   u[8] = (u[8] | 0x40) & 0x7F
//   // Set the four most significant bits (bits 12 through 15) of the
//   // time_hi_and_version field to the 4-bit version number.
//   u[6] = (u[6] & 0xF) | (0x4 << 4)
//   uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
//   return
// }

// hash plaintext with SHA-1
// func Encrypt(plaintext string) (cryptext string) {
// 	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
// 	return
// }
