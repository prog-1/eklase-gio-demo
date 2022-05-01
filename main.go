package main

import (
	"log"
	"os"

	"eklase/manager"
	"eklase/screen"

	"gioui.org/app"

	_ "modernc.org/sqlite"
)

func main() {
	// Initialize the manager.
	manager, err := manager.New("school.db")
	if err != nil {
		log.Fatal(err)
	}
	defer manager.Close()

	// Create an application UI.
	ui, err := screen.NewWindow(manager)
	if err != nil {
		log.Fatal(err)
	}

	// Run the main event loop.
	go func() {
		if err := ui.HandleEvents(manager); err != nil {
			log.Fatal(err)
		}
		// Gracefully exit the application at the end.
		os.Exit(0)
	}()
	app.Main()
}
