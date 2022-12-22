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

type ScoreBoard struct {
	score       int
	scoreText   string
	scoreWriter *text.Text
	atlas       *text.Atlas
	face        font.Face
}

func (sc *ScoreBoard) GetScoreWriter() *text.Text {
	return sc.scoreWriter
}

func (sc *ScoreBoard) Init() {
	sc.scoreText = "Score"
	sc.setFace("intuitive.ttf", 25)
	sc.setAtlas()
	sc.scoreWriter = text.New(pixel.V(50, 500), sc.atlas)
	sc.score = 0
}

func (sc *ScoreBoard) Update(dt float64, game *Game) {
	sc.scoreWriter.Clear()
	fmt.Fprintf(sc.scoreWriter, fmt.Sprintf("%s: %d", sc.scoreText, sc.score))
}

func (sc *ScoreBoard) setAtlas() {
	sc.atlas = text.NewAtlas(sc.face, text.ASCII)
}

func (sc *ScoreBoard) setFace(path string, size float64) {
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

func (sc *ScoreBoard) ChangeScore(value int) {
	sc.score = sc.score + value
}
