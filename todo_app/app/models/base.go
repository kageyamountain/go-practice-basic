package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	"todo_app/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		uuid STRING NOT NULL UNIQUE,
    		name STRING,
    		email STRING,
    		password STRING,
    		created_at DATETIME DEFAULT (datetime('now', 'localtime'))
		)`, tableNameUser)
	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		content TEXT,
    		user_id INTEGER,
    		created_at DATETIME DEFAULT (datetime('now', 'localtime'))
		)`, tableNameTodo)
	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		uuid STRING NOT NULL UNIQUE,
    		email STRING,
    		user_id INTEGER,
    		created_at DATETIME DEFAULT (datetime('now', 'localtime'))
		)`, tableNameSession)
	Db.Exec(cmdS)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj = uuid.New()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
