package bits

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func Draw(s *ebiten.Image) {

	s.Fill(color.Black)
	x := ebiten.NewImage(50, 50)
	x.Fill(color.RGBA{255, 0, 0, 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(100, 100)
	s.DrawImage(x, op)
}

