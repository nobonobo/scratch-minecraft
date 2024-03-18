package main

import (
	"embed"
	"io/fs"
	"log"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/1-16-5
var assets embed.FS

func main() {
	os.Setenv("WEBVIEW2_ADDITIONAL_BROWSER_ARGUMENTS", "--disable-web-security")
	sub, err := fs.Sub(assets, "frontend/1-16-5")
	if err != nil {
		log.Fatal(err)
	}
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options

	if err := wails.Run(&options.App{
		Title:  "scratch-minecraft",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: sub,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	}); err != nil {
		println("Error:", err.Error())
	}
}
