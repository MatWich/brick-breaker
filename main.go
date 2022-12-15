package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"github.com/MatWich/brick-breaker/classes"
)
var blocks []classes.Block = []classes.Block{}

var player = classes.Player {
	Color: colornames.Springgreen,
	Rect: pixel.R(300 - 100, 50, 300 + 100, 75),
	Vel: pixel.V(0.2, 0.2),
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
	win.SetSmooth(true)
	
	return win
}

func run() {
	win := createWindow()

	// Create stuff
	
	delimeter := 15.0
	var spaceForBlocksLastLine = 100.0
	
	var currentStart = 0.0
	for i := 0; i < 7; i++ {
		currentStart += delimeter
		blocks = append(blocks, classes.Block{
			Color: colornames.Beige,
			Rect: pixel.R(currentStart, win.Bounds().H() -100, currentStart + spaceForBlocksLastLine + delimeter, win.Bounds().H() - 50),
		})
		currentStart += spaceForBlocksLastLine + delimeter
	}

	test_ball := classes.Ball {
		Color: colornames.Green,
		Rect: pixel.C(pixel.V(win.Bounds().W(), win.Bounds().H()), 15),
		Pos: pixel.V(300, 300),
		Vel: pixel.V(0.3,0.3),
	}

	imd := imdraw.New(nil)
	imd.Precision = 32

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		win.Clear(colornames.Darkslateblue)
		
		
		// imd drawings 
		imd.Clear()
		for _, b := range blocks {
			b.Draw(imd)
		}
		player.Draw(imd)
		test_ball.Draw(imd)
		imd.Draw(win)

		// update
		player.Update(dt, *win)
		blocks = test_ball.Update(dt, win, blocks, player)
		win.Update()
	}
	
}

func main() {
	pixelgl.Run(run)
}