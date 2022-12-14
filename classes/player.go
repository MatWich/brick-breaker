package classes

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Player struct {
	Rect  pixel.Rect
	Vel   pixel.Vec
	Color color.Color

}

func (p *Player) Draw(imd *imdraw.IMDraw) {
	imd.Color = p.Color
	imd.Push(p.Rect.Min, p.Rect.Max)
	imd.Rectangle(0)
}

func (p *Player) Update(dt float64, game *Game) {
	// player movement
	if game.GetWindow().Pressed(pixelgl.KeyLeft) {
		if !(p.Rect.Min.X-p.Vel.X < 0) {
			p.Rect.Min.X -= p.Vel.X
			p.Rect.Max.X -= p.Vel.X
		}
	}

	if game.GetWindow().Pressed(pixelgl.KeyRight) {
		if !(p.Rect.Max.X+p.Vel.X > game.GetWindow().Bounds().W()) {
			p.Rect.Min.X += p.Vel.X
			p.Rect.Max.X += p.Vel.X
		}
	}

	// should not be able to go higher than 1/3 height of the screen
	if game.GetWindow().Pressed(pixelgl.KeyUp) {
		if !(p.Rect.Max.Y+p.Vel.Y > game.GetWindow().Bounds().H()/3) {
			p.Rect.Min.Y += p.Vel.Y
			p.Rect.Max.Y += p.Vel.Y
		}
	}

	if game.GetWindow().Pressed(pixelgl.KeyDown) {
		if !(p.Rect.Min.Y-p.Vel.Y < 0) {
			p.Rect.Min.Y -= p.Vel.Y
			p.Rect.Max.Y -= p.Vel.Y
		}
	}
}
