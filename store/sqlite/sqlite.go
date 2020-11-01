package sqlite

import (
	"database/sql"
	"github.com/ibreakthecloud/contact-book/store"
	"log"
	"os"

	//"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func New() *SQLite {
	filePath := "sqlite.db"
	driver := "sqlite3"

	// Remove the older DB
	os.Remove(filePath)

	//Create a new DB
	log.Println("Creating sqlite.db...")
	file, err := os.Create(filePath) // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	_ = file.Close()
	log.Println("sqlite.db created")

	sqliteDB, err := sql.Open(driver, "./"+filePath)
	if err != nil {
		panic(err)
	}

	// create contact-book table
	createTable(sqliteDB)

	return &SQLite{
		File:   filePath,
		Driver: driver,
		DB:     sqliteDB,
	}
}

func createTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE contact (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"email" TEXT,
		"name" TEXT
	);`
	statement, err := db.Prepare(createTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.Exec() // Execute SQL Statements
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("student table created")
}

func (s *SQLite) AddContact(name, email string) error {
	insertSQL := `INSERT INTO contact(email, name) VALUES (?, ?)`
	statement, err := s.DB.Prepare(insertSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		return err
	}
	_, err = statement.Exec(email, name)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLite) DeleteContact(email string) error {
	deleteSQL := `DELETE FROM contact WHERE email=?`
	statement, err := s.DB.Prepare(deleteSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		return err
	}
	_, err = statement.Exec(email)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLite) UpdateContact(email string) error {
	return nil
}

func (s *SQLite) Get(name, email string, page int) ([]store.Result, error) {
	var contacts []store.Result
	limit := 10
	offset := limit * (page - 1)

	// getAll
	row, err := s.DB.Query("SELECT * FROM contact ORDER BY name LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}

	// Below if condition will make sure if email is provided then get by email
	// or if name is provided then get by name. In case where both  are provided then,
	// this code will query by the email field, since email is unique for each contact,
	// it will return the desired result
	if email != "" {
		// getByEmail
		q := getByEmailQuery(email)
		row, err = s.DB.Query(q, email, limit, offset)
		if err != nil {
			return nil, err
		}

	} else if name != "" {
		// getByName
		q := getByNameQuery(email)
		row, err = s.DB.Query(q, name, limit, offset)
		if err != nil {
			return nil, err
		}
	}

	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var email string
		var name string
		row.Scan(&id, &email, &name)
		//log.Println("Student: ", code, " ", name, " ", program)
		newContact := store.Result{
			Id:    id,
			Name:  name,
			Email: email,
		}
		contacts = append(contacts, newContact)
	}
	return contacts, nil
}

// getByEmailQuery returns the query string to get the contact by email
func getByEmailQuery(email string) string {
	q := `SELECT * FROM contact WHERE email=? ORDER BY name LIMIT ? OFFSET ?`
	return q
}

// getByNameQuery returns the query string to get the contact by name
func getByNameQuery(name string) string {
	q := `SELECT * FROM contact WHERE name=? ORDER BY name LIMIT ? OFFSET ?`
	return q
}