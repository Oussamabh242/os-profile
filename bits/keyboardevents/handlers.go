package keyboardevents

import (
	"github.com/Oussamabh242/os-profile/bits"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

func KeysGateway(st *bits.ScreenText, kr *KeyRepeat, charMap map[ebiten.Key]string) {
	shiftHeld := ebiten.IsKeyPressed(ebiten.KeyShift) ||
		ebiten.IsKeyPressed(ebiten.KeyShiftLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyShiftRight)

	for k, val := range charMap {
		if kr.IsRepeat(k) {
			if shiftHeld {
				st.AddChar(strings.ToUpper(val))
			} else {
				st.AddChar(val)
			}

		}
	}
	if kr.IsRepeat(ebiten.KeyBackspace) {
		st.DeleteChar()
	}
	if kr.IsRepeat(ebiten.KeyEnter) {
		st.Execute()
	}
}
