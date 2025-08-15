package keyboardevents

import (
	"exmp/bits"
	"github.com/hajimehoshi/ebiten/v2"
)

func KeysGateway(st *bits.ScreenText, kr *KeyRepeat, charMap map[ebiten.Key]string) {
	for k, val := range charMap {
		if kr.IsRepeat(k) {
			st.AddChar(val)
		}
	}
	if kr.IsRepeat(ebiten.KeyBackspace) {
		st.DeleteChar()
	}
	if kr.IsRepeat(ebiten.KeyEnter) {
		st.Execute()
	}
}
