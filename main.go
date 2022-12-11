package main

import (
	"image/color"
	"time"

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

func (b * ball) update(dt float64, win *pixelgl.Window, blocks block) {
	b.pos.X += b.vel.X
	b.pos.Y += b.vel.Y


	// collision with walls
	if b.pos.X - b.rect.Radius < 0 || b.pos.X + b.rect.Radius > win.Bounds().W() {
		b.vel.X *= -1
	}

	if b.pos.Y - b.rect.Radius < 0 || b.pos.Y +b.rect.Radius > win.Bounds().H() {
		b.vel.Y *= -1
	}

	// colision with block

	if (b.rect.IntersectRect(blocks.rect) != pixel.V(-0, -0)) {
		if (b.pos.X + b.rect.Radius >= blocks.rect.Min.X && b.pos.X - b.rect.Radius <= blocks.rect.Max.X) {
			b.vel.X *= -1
			b.pos.X += b.vel.X
		}
	
		if (b.pos.Y + b.rect.Radius >= blocks.rect.Min.Y && b.pos.Y - b.rect.Radius <= blocks.rect.Max.Y) {
			b.vel.Y *= -1
			b.pos.Y += b.vel.Y
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
	test_block := block {
		color: colornames.Red,
		rect: pixel.R(5, win.Bounds().H() -100, 205, win.Bounds().H() -10),
	}

	test_ball := ball {
		color: colornames.Green,
		rect: pixel.C(pixel.V(win.Bounds().W(), win.Bounds().H()), 15),
		pos: pixel.V(300, 300),
		vel: pixel.V(0.3,0.3),
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
		test_block.draw(imd)
		test_ball.draw(imd)
		imd.Draw(win)

		// update
		test_ball.update(dt, win, test_block)
		win.Update()
	}
	
}

func main() {
	pixelgl.Run(run)
}