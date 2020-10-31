package sqlite

import "database/sql"

type SQLite struct {
	DB      *sql.DB
	File    string
	Driver string
}
