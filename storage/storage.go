// Package storage is an interface for interacting with a database.
package storage

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var (
	// Statement for creating tables. Currently creates `students` table only.
	// In the future can be expanded with creation of other tables.
	createTableStmt = `
CREATE TABLE IF NOT EXISTS students (
	id	    INTEGER PRIMARY KEY AUTOINCREMENT,
	name	  TEXT,
	surname	TEXT
);`
	// Statement for adding a new entry into `students` table.
	insertStudentsStmt = `INSERT INTO students (name, surname) VALUES(?, ?)`
	// Statement for getting all entries from `students` table.
	selectStudentsStmt = `SELECT * FROM students`
)

// StudentEntry represents a row for a single student in the DB.
type StudentEntry struct {
	ID      int64
	Name    string `db:"name"`
	Surname string `db:"surname"`
}

// Storage is an interface for interacting with persistent storage.
type Storage struct {
	db *sqlx.DB
}

// Open initializes a new DB given its path, or opens an existing DB, and
// initializes the handler. Returns an error if any of the steps fails.
func Open(path string) (s *Storage, err error) {
	var db *sqlx.DB
	// Open a DB by the path.
	if db, err = sqlx.Open("sqlite", path); err != nil {
		return nil, err
	}
	// Create new tables. Note that the tables may exist already.
	if _, err = db.Exec(createTableStmt); err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func MustOpen(path string) *Storage {
	s, err := Open(path)
	if err != nil {
		log.Fatalf("unable to create storage: %v", err)
	}
	return s
}

// Close closes the database after it is no longer required.
func (s *Storage) Close() error { return s.db.Close() }

// Students returns a slice of existing students.
func (s Storage) Students() (entries []StudentEntry, err error) {
	err = s.db.Select(&entries, selectStudentsStmt)
	return entries, err
}

// AddStudent appends a new student entry to the database.
func (s *Storage) AddStudent(name, surname string) error {
	// Attempt to add an entry to the database first.
	// If it fails, the student field will not be modified.
	_, err := s.db.Exec(insertStudentsStmt, name, surname)
	return err
}
