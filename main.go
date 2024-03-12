package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	// Initialize the application
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create a main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Qt Window")

	// Set the main window size
	window.SetMinimumSize2(400, 300)

	// Show the main window
	window.Show()

	// Start the application event loop
	app.Exec()
}
