package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/skip2/go-qrcode"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"golang.org/x/sys/windows/registry"
)

const (
	regKeyPath    = `Software\BioSecQRID`
	regValueName  = "PlainText"
	qrCodeFile    = "output_qr_code.png"
	sleepDuration = 30 * time.Second
)

func main() {
	// Read the device ID from the registry
	deviceID, err := GetDeviceIDFromRegistry()
	if err != nil {
		log.Fatal("Error reading device ID from registry:", err)
	}

	// Generate the QR code
	err = GenerateQRCode(deviceID, qrCodeFile)
	if err != nil {
		log.Fatal("Error generating QR code:", err)
	}
	fmt.Printf("QR code for Device ID '%s' has been generated and saved as '%s'.\n", deviceID, qrCodeFile)

	// Create and run the Qt application
	widgetApp := widgets.NewQApplication(len(os.Args), os.Args)

	// Create a main window
	mainWindow := widgets.NewQMainWindow(nil, 0)
	mainWindow.SetWindowTitle("QR Code Viewer")

	// Remove the title bar
	mainWindow.SetWindowFlags(core.Qt__FramelessWindowHint)

	// Create a QLabel to display the image
	imageLabel := widgets.NewQLabel2("", nil, 0)
	pixmap := gui.NewQPixmap()
	pixmap.Load(qrCodeFile, "", core.Qt__AutoColor)
	imageLabel.SetPixmap(pixmap)

	// Set up the main window layout
	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(imageLabel, 0, 0)
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)
	mainWindow.SetCentralWidget(widget)

	// Show the main window
	mainWindow.Show()

	// Sleep for a specified duration
	time.Sleep(sleepDuration)

	// Delete the QR code image
	deleteQRCodeFile(qrCodeFile)

	// Run the application event loop
	widgetApp.Exec()
}

func deleteQRCodeFile(filename string) {
	// Delete the QR code file
	err := os.Remove(filename)
	if err != nil {
		fmt.Println("Error deleting QR code file:", err)
	}
}

func GetDeviceIDFromRegistry() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, regKeyPath, registry.READ)
	if err != nil {
		return "", err
	}
	defer k.Close()

	deviceID, _, err := k.GetStringValue(regValueName)
	if err != nil {
		return "", err
	}

	return deviceID, nil
}

func GenerateQRCode(deviceID, filename string) error {
	return qrcode.WriteFile(deviceID, qrcode.Low, 256, filename)
}
