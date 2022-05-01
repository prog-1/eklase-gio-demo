package main

import (
	"log"
	"os"

	"eklase/screen"
	"eklase/state"

	"gioui.org/app"

	_ "modernc.org/sqlite"
)

func main() {
	// Initialize the state.
	state, err := state.New("school.db")
	if err != nil {
		log.Fatal(err)
	}
	defer state.Close()

	// Create an application UI.
	ui, err := screen.NewHandle(state)
	if err != nil {
		log.Fatal(err)
	}

	// Run the main event loop.
	go func() {
		if err := ui.HandleEvents(); err != nil {
			log.Fatal(err)
		}
		// Gracefully exit the application at the end.
		os.Exit(0)
	}()
	app.Main()
}
