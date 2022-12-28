package classes

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type HUD struct {
	score       int
	scoreText   string
	scoreWriter *text.Text
	atlas       *text.Atlas
	face        font.Face
	lives       int
	livesText   string
	livesWriter *text.Text
}

func (hud *HUD) GetScoreWriter() *text.Text {
	return hud.scoreWriter
}

func (hud *HUD) GetLivesWritter() *text.Text {
	return hud.livesWriter
}

func (hud *HUD) GetLives() int {
	return hud.lives
}

func (hud *HUD) Init() {
	// Score init
	hud.scoreText = "Score"
	hud.setFace("intuitive.ttf", 25)
	hud.setAtlas()
	hud.scoreWriter = text.New(pixel.V(50, 500), hud.atlas)
	hud.score = 0
	// Lives init
	hud.lives = 3
	hud.livesText = "Lives"
	hud.livesWriter = text.New(pixel.V(920, 740), hud.atlas)
}

func (hud *HUD) Update(dt float64, game *Game) {
	hud.scoreWriter.Clear()
	hud.livesWriter.Clear()
	fmt.Fprintf(hud.scoreWriter, fmt.Sprintf("%s: %d", hud.scoreText, hud.score))
	fmt.Fprintf(hud.livesWriter, fmt.Sprintf("%s: %d", hud.livesText, hud.lives))
}

func (hud *HUD) setAtlas() {
	hud.atlas = text.NewAtlas(hud.face, text.ASCII)
}

func (hud *HUD) setFace(path string, size float64) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return
	}

	hud.face = truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	})
}

func (hud *HUD) ChangeScore(value int) {
	hud.score += value
}

func (hud *HUD) ChangeLives(value int) {
	hud.lives += value
}

func (hud *HUD) Reset() {
	hud.lives = 3
	hud.score = 0
}