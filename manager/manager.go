// Package manager stores the application state and and provides accessors to the
// storage.
package manager

import (
	"eklase/storage"
)

// Handle is a state handler.
type AppManager struct {
	// Database handler.
	storage *storage.Storage
	// Specifies if quitting the application was requested.
	quit bool
}

// New returns a new state handler. Returns an error if any of the steps fails.
func New(path string) (*AppManager, error) {
	storage, err := storage.New(path)
	if err != nil {
		return nil, err
	}
	return &AppManager{storage: storage}, nil
}

// Close closes the dependencies when they are no longer required.
func (h *AppManager) Close() error {
	return h.storage.Close()
}

// StudentCount returns the total number of students.
func (h *AppManager) StudentCount() (int, error) {
	rows, err := h.storage.Students()
	if err != nil {
		return -1, err
	}
	return len(rows), nil
}

// Student accesses a student by its index.
//
// The index is not guaranteed to be equal to the ID column value in the
// database.
func (h *AppManager) Student(i int) (storage.StudentEntry, error) {
	entries, err := h.storage.Students()
	if err != nil {
		return storage.StudentEntry{}, err
	}
	return entries[i], nil
}

// AddStudent adds a student to the database.
func (v *AppManager) AddStudent(name, surname string) {
	v.storage.AddStudent(name, surname)
}

// Quit requests quitting the application.
func (v *AppManager) Quit() {
	v.quit = true
}

// ShouldQuit returns whether quitting the application was requested.
func (v *AppManager) ShouldQuit() bool {
	return v.quit
}
