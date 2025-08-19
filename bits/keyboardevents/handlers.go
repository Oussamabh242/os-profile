package keyboardevents

import (
	"fmt"
	"strings"

	"github.com/Oussamabh242/os-profile/context"

	"github.com/Oussamabh242/os-profile/bits"

	"github.com/hajimehoshi/ebiten/v2"
)

func KeysGateway(st *bits.ScreenText, kr *KeyRepeat, charMap map[ebiten.Key]string, controller *uint8) {
	shiftHeld := ebiten.IsKeyPressed(ebiten.KeyShift) ||
		ebiten.IsKeyPressed(ebiten.KeyShiftLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyShiftRight)

	xoff, yoff := ebiten.Wheel()
	if xoff != 0 || yoff != 0 {

		fmt.Printf("mouse scrlled x : %f ,,  y : %f \n ", xoff, yoff)
	}

	for k, val := range charMap {
		if kr.IsRepeat(k) {
			if shiftHeld {
				st.AddChar(strings.ToUpper(val))
			} else {
				if *controller == 1 {
					if val == "q" {
						*controller = 0
						context.CatStartY = 10
						return
					}
					if val == "n" {
						context.CatStartY -= 30
						return
					}
					if val == "p" {
						context.CatStartY += 30
						return
					}
				}
				// if yoff != 0 && *controller == 1 {
				// 	context.CatStartY += int(yoff)
				// }
				st.AddChar(val)
			}

		}
	}
	if kr.IsRepeat(ebiten.KeyBackspace) {
		st.DeleteChar()
	}
	if kr.IsRepeat(ebiten.KeyEnter) {
		st.Execute(controller)
	}
}
