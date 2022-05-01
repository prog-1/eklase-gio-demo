// Package state stores the application state and and provides accessors to the
// database.
package state

import "eklase/db"

// Handle is a state handler.
type Handle struct {
	// Database handler.
	dbHandle *db.Handle
	// Specifies if quitting the application was requested.
	quit bool
}

// New returns a new state handler. Returns an error if any of the steps fails.
func New(path string) (*Handle, error) {
	db, err := db.New(path)
	if err != nil {
		return nil, err
	}
	return &Handle{dbHandle: db}, nil
}

// Close closes the dependencies when they are no longer required.
func (h *Handle) Close() error {
	return h.dbHandle.Close()
}

// StudentCount returns the total number of students.
func (h *Handle) StudentCount() int {
	return len(h.dbHandle.Students())
}

// Student accesses a student by its index.
//
// The index is not guaranteed to be equal to the ID column value in the
// database.
func (h *Handle) Student(i int) db.StudentEntry {
	return h.dbHandle.Students()[i]
}

// AddStudent adds a student to the database.
func (v *Handle) AddStudent(name, surname string) {
	v.dbHandle.AddStudent(name, surname)
}

// Quit requests quitting the application.
func (v *Handle) Quit() {
	v.quit = true
}

// ShouldQuit returns whether quitting the application was requested.
func (v *Handle) ShouldQuit() bool {
	return v.quit
}
