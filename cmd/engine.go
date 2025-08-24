package main

import (
	"strings"

	"github.com/Oussamabh242/os-profile/bits"
	"github.com/Oussamabh242/os-profile/bits/keyboardevents"
	"github.com/Oussamabh242/os-profile/context"

	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// var openURL = js.Global().Get("openURL")

var MdCache map[string]string = make(map[string]string)

var lastUpdate time.Time
var kr = keyboardevents.NewKeyRepeat()

const (
	SCREEN_WIDTH  = 1920
	SCREEN_HEIGHT = 1080
)

type Game struct {
	// visibleLines int
	ScreenText         *bits.ScreenText
	ticks              int
	Controller         uint8 // 0 -> tty , // 1 -> cat cmd
	FileHoldingControl string
}

func (g *Game) Update() error {
	kr.Update()

	g.ticks++
	g.ScreenText.Update(g.ticks)

	// if kr.IsRepeat(ebiten.KeyControl) {
	// openURL.Invoke("https://example.com")
	// }

	keyboardevents.KeysGateway(g.ScreenText, kr, context.KeyToChar, &g.Controller, &g.FileHoldingControl)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if g.Controller == 1 {
		// data, err := static.PostsFS.ReadFile(g.FileHoldingControl)
		splitted := strings.Split(g.FileHoldingControl, ".")
		if len(splitted) != 2 {
			log.Println("oops cannot splitt perfectly")
			g.Controller = 0
		}
		v, exists := MdCache[splitted[0]]
		if !exists {
			str := context.MakeOutsideReqeust("/posts/" + splitted[0])
			if str == "" {
				log.Println("returned emptyFile")
				g.Controller = 0
			}

			MdCache[splitted[0]] = str
		}

		bits.DrawBase(screen, v)

		return
	}

	g.ScreenText.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	ebiten.SetTPS(30)
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {

	st := bits.NewScreenText()
	//Game Instance
	game := &Game{
		ScreenText: st,
		Controller: 0,
	}
	// rant-on-typescript.

	//startmenu init

	// ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)

	ebiten.SetTPS(30)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

//// WASM
