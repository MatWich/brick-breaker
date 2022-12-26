package main

import (
	"time"

	"github.com/MatWich/brick-breaker/classes"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var game = classes.Game{}

func run() {

	// Create stuff
	game.CreateWindow()
	game.CreateBlocks()
	game.CreateBall()
	game.CreatePlayer()
	game.CreateScoreBoard()

	imd := imdraw.New(nil)
	imd.Precision = 32

	// Main loop
	last := time.Now()
	for !game.GetWindow().Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		game.GetWindow().Clear(colornames.Darkslateblue)

		// imd drawings
		imd.Clear()
		for _, b := range game.GetBlocks() {
			b.Draw(imd)
		}
		game.GetPlayer().Draw(imd)
		game.GetBall().Draw(imd)
		game.GetHUD().GetScoreWriter().Draw(game.GetWindow(), pixel.IM.Moved(pixel.V(-40, 240)))
		game.GetHUD().GetLivesWritter().Draw(game.GetWindow(), pixel.IM)
		imd.Draw(game.Window)

		// update
		game.GetPlayer().Update(dt, &game)
		game.GetBall().Update(dt, &game)
		game.GetHUD().Update(dt, &game)
		game.GetWindow().Update()
	}

}

func main() {
	pixelgl.Run(run)
}
