package classes

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Game struct {
	player     Player
	Window     *pixelgl.Window
	ball       Ball
	blocks     []Block
	hud HUD
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

func (g *Game) GetHUD() *HUD {
	return &g.hud
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
	g.blocks = make([]Block, 0, 21)
	delimeter := 15.0

	window_width := g.GetWindow().Bounds().W()
	space_for_block := window_width - delimeter*8
	space_per_block := space_for_block / 7

	window_height := g.GetWindow().Bounds().H()
	block_height := 50.0

	var currentStart = 0.0
	for j := 0; j > -3; j-- {
		currentStart = delimeter
		for i := 0; i < 7; i++ {
			g.blocks = append(g.blocks, Block{
				Color: colornames.Azure,
				Rect:  pixel.R(currentStart, window_height-100+float64(j)*(block_height+delimeter), currentStart+space_per_block, window_height-50+float64(j)*(block_height+delimeter)),
			})
			currentStart += space_per_block + delimeter
		}

	}
}

func (g *Game) CreateBall() {
	g.ball = Ball{
		Color: colornames.Green,
		Rect:  pixel.C(pixel.V(g.GetWindow().Bounds().W(), g.GetWindow().Bounds().H()), 15),
		Pos:   pixel.V(300, 300),
		Vel:   pixel.V(0.3, 0.3),
	}
}

func (g *Game) CreatePlayer() {
	g.player = Player{
		Color: colornames.Springgreen,
		Rect:  pixel.R(300-100, 50, 300+100, 75),
		Vel:   pixel.V(0.3, 0.3),
	}
}

func (g *Game) CreateScoreBoard() {
	g.hud = HUD{}
	g.hud.Init()
}


func (g *Game) Reset(full bool) {
	g.CreateBall()
	g.CreateBlocks()
	if full {
		g.hud.Reset()
		g.CreatePlayer()
	}
}