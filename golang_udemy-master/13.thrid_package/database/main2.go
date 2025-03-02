package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// SQLiteのデータベースファイルを開く（または作成）
	DbConnection, err := sql.Open("sqlite3", "./example.sql")
	if err != nil {
		log.Fatalf("DB接続失敗: %v", err)
	}
	defer DbConnection.Close() // 関数終了時にDBを閉じる

	// SQL実行
	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	_, err = DbConnection.Exec(cmd, "Nancy", 20)
	if err != nil {
		log.Fatalf("SQL実行失敗: %v", err)
	}

	log.Println("データが挿入されました")
}
