package sqlite

import (
	"database/sql"
	"log"
	"os"
)

func New() *SQLite {
	filePath := "sqlite.db"
	driver := "sqlite3"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// file does not exist
		log.Println("Creating sqlite.db...")
		file, err := os.Create(filePath)// Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		_ = file.Close()
		log.Println("sqlite.db created")
	}
	sqliteDB, err := sql.Open(driver, "/." +filePath)
	if err != nil {
		panic(err)
	}

	return &SQLite{
		File:    filePath,
		Driver: driver,
		DB: sqliteDB,
	}
}
