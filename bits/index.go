package bits

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable interface {
	Draw(screen *ebiten.Image)
}

type Bit struct {
	Color color.Color
	Img   *ebiten.Image
	Ops   *ebiten.DrawImageOptions
	Yes   bool
}

func NewBit(c color.Color, H, W, SX, SY float64) *Bit {

	op := &ebiten.DrawImageOptions{}
	// op.GeoM.Scale(W, H)
	op.GeoM.Translate(SX, SY)

	img := ebiten.NewImage(int(W), int(H))
	img.Fill(c)

	return &Bit{
		Color: c,
		Ops:   op,
		Img:   img,
		Yes:   true,
	}
}

func (b *Bit) Draw(screen *ebiten.Image) {
	if b.Yes {

		screen.DrawImage(b.Img, b.Ops)
	}

}
