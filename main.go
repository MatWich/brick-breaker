package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func createWindow() *pixelgl.Window {
	cfg :=pixelgl.WindowConfig{
		Title: "Brick Breaker",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync: false,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	return win
}

func run() {
	win := createWindow()

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}