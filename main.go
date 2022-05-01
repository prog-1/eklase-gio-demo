package main

import (
	"log"
	"os"

	"eklase/screen"
	"eklase/state"
	"eklase/storage"

	"gioui.org/app"

	_ "modernc.org/sqlite"
)

func main() {
	storage := storage.Must(storage.New("school.db")) // We'll defer Close later.
	state := state.New(storage)

	// Run the main event loop.
	go func() {
		ui, err := screen.NewWindow(state)
		if err != nil {
			log.Fatal(err)
		}

		func() { // For deferred calls.
			defer storage.Close()
			err = ui.HandleEvents(state)
		}()
		if err != nil {
			log.Fatalf("failed to handle events: %v", err)
		}
		// Gracefully exit the application at the end.
		os.Exit(0)
	}()
	app.Main()
}
