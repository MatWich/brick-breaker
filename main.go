package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type block struct {
	rect pixel.Rect
	color color.Color
}

func (b *block) draw(imd *imdraw.IMDraw) {
	imd.Color = b.color
	imd.Push(b.rect.Min, b.rect.Max)
	imd.Rectangle(0)
}

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

	// Create stuff
	test_block := block {
		color: colornames.Red,
		rect: pixel.R(5, win.Bounds().H() -100, 105, win.Bounds().H() -50),
	}

	imd := imdraw.New(nil)
	imd.Precision = 32

	for !win.Closed() {
		win.Clear(colornames.Darkslateblue)
		test_block.draw(imd)
		
		imd.Draw(win)
		win.Update()
	}
	
}

func main() {
	pixelgl.Run(run)
}