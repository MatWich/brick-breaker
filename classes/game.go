package classes

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Game struct {
	player Player
	Window *pixelgl.Window
	ball   Ball
	blocks []Block
	ScoreBoard ScoreBoard
}

// GETTERS AND SETTERS
func (g *Game) GetWindow() *pixelgl.Window {
	return g.Window
}

func (g *Game) GetBlocks() []Block {
	return g.blocks
}

func (g *Game) SetBlocks(blocks []Block) {
	g.blocks = blocks
}

func (g *Game) GetPlayer() *Player {
	return &g.player
}

func (g *Game) GetBall() *Ball {
	return &g.ball
}

// INITIALIZATION

func (g *Game) CreateWindow() {
	cfg := pixelgl.WindowConfig{
		Title:  "Brick Breaker",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  false,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	g.Window = win
}

func (g *Game) CreateBlocks() {
	delimeter := 15.0
	var spaceForBlocksLastLine = 100.0

	var currentStart = 0.0
	for i := 0; i < 7; i++ {
		currentStart += delimeter
		g.blocks = append(g.blocks, Block{
			Color: colornames.Azure,
			Rect:  pixel.R(currentStart, g.Window.Bounds().H()-100, currentStart+spaceForBlocksLastLine+delimeter, g.Window.Bounds().H()-50),
		})
		currentStart += spaceForBlocksLastLine + delimeter
	}
}

func (g *Game) CreateBall() {
	g.ball = Ball{
		Color: colornames.Green,
		Rect:  pixel.C(pixel.V(g.Window.Bounds().W(), g.Window.Bounds().H()), 15),
		Pos:   pixel.V(300, 300),
		Vel:   pixel.V(0.3, 0.3),
	}
}

func (g *Game) CreatePlayer() {
	g.player = Player{
		Color: colornames.Springgreen,
		Rect:  pixel.R(300-100, 50, 300+100, 75),
		Vel:   pixel.V(0.2, 0.2),
	}
}

func (g *Game) CreateScoreBoard() {
	g.ScoreBoard = ScoreBoard{}
	g.ScoreBoard.Init()
}