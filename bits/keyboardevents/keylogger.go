package keyboardevents

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type KeyRepeat struct {
	heldFrames   map[ebiten.Key]int
	initialDelay int
	repeatRate   int
}

func NewKeyRepeat() *KeyRepeat {
	return &KeyRepeat{
		heldFrames:   make(map[ebiten.Key]int),
		initialDelay: 10, // frames before first repeat
		repeatRate:   1,  // frames between repeats
	}
}

func (kr *KeyRepeat) Update() {
	for k := range kr.heldFrames {
		if ebiten.IsKeyPressed(k) {
			kr.heldFrames[k]++
		} else {
			delete(kr.heldFrames, k)
		}
	}
}

func (kr *KeyRepeat) IsRepeat(k ebiten.Key) bool {
	if inpututil.IsKeyJustPressed(k) {
		kr.heldFrames[k] = 0
		return true
	}
	if frames, ok := kr.heldFrames[k]; ok {
		if frames >= kr.initialDelay && (frames-kr.initialDelay)%kr.repeatRate == 0 {
			return true
		}
	}
	return false
}
