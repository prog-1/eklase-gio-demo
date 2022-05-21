package state

import (
	"eklase/storage"
)

// State is the application context (aka state). It provides access to the
// features that do not depend on implementation e.g. (T)UI framework.
type State struct {
	storage *storage.Storage // Provides DB access.

	quit bool // True if the application should exit.
}

// New returns a new state handler. Returns an error if any of the steps fails.
func New(s *storage.Storage) *State { return &State{storage: s} }

// Students returns students stored in the database.
func (h *State) Students() ([]storage.StudentEntry, error) {
	return h.storage.Students()
}

// AddStudent adds a student to the database.
func (v *State) AddStudent(name, surname string) error {
	return v.storage.AddStudent(name, surname)
}

// Quit requests quitting the application.
func (v *State) Quit() { v.quit = true }

// ShouldQuit returns whether quitting the application was requested. The
// method does not reset internal quit flag and keeps it set.
func (v *State) ShouldQuit() bool { return v.quit }
