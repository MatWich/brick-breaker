package classes

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Ball struct {
	Rect pixel.Circle
	Color color.Color
	Pos pixel.Vec
	Vel pixel.Vec
}

func (b *Ball) Draw(imd *imdraw.IMDraw) {
	imd.Color = b.Color
	imd.Push(b.Pos)
	imd.Circle(15, 0)
}

func (b * Ball) Update(dt float64, win *pixelgl.Window, blocks []Block, player Player) []Block {
	b.Pos.X += b.Vel.X
	b.Pos.Y += b.Vel.Y


	// collision with walls
	if b.Pos.X - b.Rect.Radius < 0 || b.Pos.X + b.Rect.Radius > win.Bounds().W() {
		b.Vel.X *= -1
	}

	if b.Pos.Y - b.Rect.Radius < 0 || b.Pos.Y + b.Rect.Radius > win.Bounds().H() {
		b.Vel.Y *= -1
	}

	// colision with block
	toDelete := []int{}
	for i, blk := range blocks {
		collision := b.Rect.IntersectRect(blk.Rect)
		 
		if (b.Rect.IntersectRect(blk.Rect) != pixel.V(-0, -0)) {
			
			if math.Abs(collision.Y) > math.Abs(collision.X) {
				b.Vel.Y *= -1
				toDelete = append(toDelete, i)

			} else {
				b.Vel.X *= -1
				toDelete = append(toDelete, i)
			}
		}
		
	}

	for _, i := range toDelete {
		blocks = append(blocks[:i], blocks[i+1:]...)
	}

	// colision with player
	var collision = b.Rect.IntersectRect(player.Rect)
	if collision != pixel.V(-0, -0) {
		if math.Abs(collision.Y) > math.Abs(collision.X) {
			b.Vel.Y *= -1
			b.Pos.Y += 2 * b.Vel.Y +collision.Y
		} else {
			b.Vel.X *= -1
			b.Pos.X += 2 * b.Vel.X + collision.X
		}
	} 
		
	b.Rect.Center.X = b.Pos.X
	b.Rect.Center.Y = b.Pos.Y

	return blocks
}