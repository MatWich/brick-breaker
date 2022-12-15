package classes

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Block struct {
	Rect pixel.Rect
	Color color.Color
}

func (b *Block) Draw(imd *imdraw.IMDraw) {
	imd.Color = b.Color
	imd.Push(b.Rect.Min, b.Rect.Max)
	imd.Rectangle(0)
}