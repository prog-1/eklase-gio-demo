// Package db is an interface for interacting with a database.
package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

var (
	// Statement for creating tables. Currently creates `students` table only.
	// In the future can be expanded with creation of other tables.
	createTableStmt = `
CREATE TABLE IF NOT EXISTS students (
	id	INTEGER,
	name	TEXT,
	surname	TEXT,
	PRIMARY KEY(id AUTOINCREMENT)
);`
	// Statement for adding a new entry into `students` table.
	insertStudentsStmt = `INSERT INTO students (name, surname) VALUES(?, ?)`
	// Statement for getting all entries from `students` table.
	selectStudentsStmt = `SELECT name, surname FROM students`
)

// StudentEntry represents a row for a single student in the DB.
type StudentEntry struct {
	Name    string `db:"name"`
	Surname string `db:"surname"`
}

// Handle is a DB handler.
//
// Handle ensures that its fields are always consistent with the database state.
type Handle struct {
	db       *sqlx.DB
	students []StudentEntry
}

// New initializes a new DB given its path, or opens an existing DB, and
// initializes the handler. Returns an error if any of the steps fails.
func New(path string) (*Handle, error) {
	// Open a DB by the path.
	db, err := sqlx.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite DB: %v", db)
	}

	// Create new tables. Note that the tables may exist already.
	res, err := db.Exec(createTableStmt)
	if err != nil {
		return nil, fmt.Errorf("table creation failed. Query: %v\nError: %v", createTableStmt, err)
	}
	if cnt, err := res.RowsAffected(); err != nil {
		log.Printf("%d rows affected.", cnt)
	}
	h := Handle{db: db}

	// Read rows from the `students` table and populate students field in the
	// handler.
	if err := h.db.Select(&h.students, selectStudentsStmt); err != nil {
		return nil, fmt.Errorf("querying 'students' table failed. Query: %v\nError: %v", selectStudentsStmt, err)
	}

	return &h, nil
}

// Close closes the database after it is no longer required.
func (h *Handle) Close() error {
	return h.db.Close()
}

// Students returns a slice of existing students.
func (h Handle) Students() []StudentEntry {
	return h.students
}

// AddStudent appends a new student entry to the database.
func (h *Handle) AddStudent(name, surname string) error {
	// Attempt to add an entry to the database first.
	// If it fails, the student field will not be modified.
	res, err := h.db.Exec(insertStudentsStmt, name, surname)
	if err != nil {
		return fmt.Errorf("table creation failed. Query: %v\nError: %v", createTableStmt, err)
	}
	if cnt, err := res.RowsAffected(); err != nil {
		log.Printf("%d rows affected.", cnt)
	}

	h.students = append(h.students, StudentEntry{name, surname})

	return nil
}
