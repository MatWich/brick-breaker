package main

import (
	"image/color"
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)
var blocks []block = make([]block, 7)

type player struct {
	rect pixel.Rect
	vel pixel.Vec
	color color.Color
}

func (p *player) draw(imd *imdraw.IMDraw) {
	imd.Color = p.color
	imd.Push(p.rect.Min, p.rect.Max)
	imd.Rectangle(0)
}

func (p *player) update(dt float64, win pixelgl.Window) {
	// player movement
	if win.Pressed(pixelgl.KeyLeft) {
		p.rect.Min.X -= p.vel.X
		p.rect.Max.X -= p.vel.X
	}

	if win.Pressed(pixelgl.KeyRight) {
		p.rect.Min.X += p.vel.X
		p.rect.Max.X += p.vel.X
	}

	if win.Pressed(pixelgl.KeyUp) {
		p.rect.Min.Y += p.vel.Y
		p.rect.Max.Y += p.vel.Y
	}

	if win.Pressed(pixelgl.KeyDown) {
		p.rect.Min.Y -= p.vel.Y
		p.rect.Max.Y -= p.vel.Y
	}


}

type block struct {
	rect pixel.Rect
	color color.Color
}

func (b *block) draw(imd *imdraw.IMDraw) {
	imd.Color = b.color
	imd.Push(b.rect.Min, b.rect.Max)
	imd.Rectangle(0)
}

type ball struct {
	rect pixel.Circle
	color color.Color
	pos pixel.Vec
	vel pixel.Vec
}

func (b *ball) draw(imd *imdraw.IMDraw) {
	imd.Color = b.color
	imd.Push(b.pos)
	imd.Circle(15, 0)
}

func (b * ball) update(dt float64, win *pixelgl.Window) {
	b.pos.X += b.vel.X
	b.pos.Y += b.vel.Y


	// collision with walls
	if b.pos.X - b.rect.Radius < 0 || b.pos.X + b.rect.Radius > win.Bounds().W() {
		b.vel.X *= -1
	}

	if b.pos.Y - b.rect.Radius < 0 || b.pos.Y + b.rect.Radius > win.Bounds().H() {
		b.vel.Y *= -1
	}

	// colision with block
	for i, blk := range blocks {
		collision := b.rect.IntersectRect(blk.rect)
		 
		if (b.rect.IntersectRect(blk.rect) != pixel.V(-0, -0)) {
			
			// if (b.pos.Y + b.rect.Radius >= blk.rect.Min.Y && b.pos.Y - b.rect.Radius <= blk.rect.Max.Y) {
			if math.Abs(collision.Y) > math.Abs(collision.X) {
				blocks = append(blocks[:i], blocks[i+1:]...)
				b.vel.Y *= -1
				continue
			} else {
				blocks = append(blocks[:i], blocks[i+1:]...)
				b.vel.X *= -1
			}
		}
	}

	b.rect.Center.X = b.pos.X
	b.rect.Center.Y = b.pos.Y
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
	for i := range blocks {
		currentStart += delimeter
		blocks[i] = block{
			color: colornames.Beige,
			rect: pixel.R(currentStart, win.Bounds().H() -100, currentStart + spaceForBlocksLastLine + delimeter, win.Bounds().H() - 50),
		}
		currentStart += spaceForBlocksLastLine + delimeter
	}

	test_ball := ball {
		color: colornames.Green,
		rect: pixel.C(pixel.V(win.Bounds().W(), win.Bounds().H()), 15),
		pos: pixel.V(300, 300),
		vel: pixel.V(0.1,0.3),
	}

	test_player := player {
		color: colornames.Springgreen,
		rect: pixel.R(300 - 100, 50, 300 + 100, 75),
		vel: pixel.V(0.2, 0.2),
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
			b.draw(imd)
		}
		test_player.draw(imd)
		test_ball.draw(imd)
		imd.Draw(win)

		// update
		test_player.update(dt, *win)
		test_ball.update(dt, win)
		win.Update()
	}
	
}

func main() {
	pixelgl.Run(run)
}