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

func (sc *HUD) GetScoreWriter() *text.Text {
	return sc.scoreWriter
}

func (sc *HUD) GetLivesWritter() *text.Text {
	return sc.livesWriter
}

func (sc *HUD) GetLives() int {
	return sc.lives
}

func (sc *HUD) Init() {
	// Score init
	sc.scoreText = "Score"
	sc.setFace("intuitive.ttf", 25)
	sc.setAtlas()
	sc.scoreWriter = text.New(pixel.V(50, 500), sc.atlas)
	sc.score = 0
	// Lives init
	sc.lives = 3
	sc.livesText = "Lives"
	sc.livesWriter = text.New(pixel.V(920, 740), sc.atlas)
}

func (sc *HUD) Update(dt float64, game *Game) {
	sc.scoreWriter.Clear()
	sc.livesWriter.Clear()
	fmt.Fprintf(sc.scoreWriter, fmt.Sprintf("%s: %d", sc.scoreText, sc.score))
	fmt.Fprintf(sc.livesWriter, fmt.Sprintf("%s: %d", sc.livesText, sc.lives))
}

func (sc *HUD) setAtlas() {
	sc.atlas = text.NewAtlas(sc.face, text.ASCII)
}

func (sc *HUD) setFace(path string, size float64) {
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

	sc.face = truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	})
}

func (sc *HUD) ChangeScore(value int) {
	sc.score += value
}

func (sc *HUD) ChangeLives(value int) {
	sc.lives += value
}