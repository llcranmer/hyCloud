package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "World!"
}

func main() {
	js := mewn.String("./frontend/public/build/bundle.js")
	css := mewn.String("./frontend/public/build/bundle.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "hycloud",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(basic)
	app.Run()
}
