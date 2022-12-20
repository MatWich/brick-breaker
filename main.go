package main

import (
	"github.com/MatWich/brick-breaker/classes"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

var game = classes.Game{}

func run() {

	// Create stuff

	game.CreateWindow()
	game.CreateBlocks()
	game.CreateBall()
	game.CreatePlayer()

	imd := imdraw.New(nil)
	imd.Precision = 32

	// Main loop
	last := time.Now()
	for !game.Window.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		game.Window.Clear(colornames.Darkslateblue)

		// imd drawings
		imd.Clear()
		for _, b := range game.GetBlocks() {
			b.Draw(imd)
		}
		game.GetPlayer().Draw(imd)
		game.GetBall().Draw(imd)
		imd.Draw(game.Window)

		// update
		game.GetPlayer().Update(dt, *game.Window)
		game.SetBlocks(game.GetBall().Update(dt, game.Window, game.GetBlocks(), game.GetPlayer()))
		game.Window.Update()
	}

}

func main() {
	pixelgl.Run(run)
}
