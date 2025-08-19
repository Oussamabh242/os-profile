package main

import (
	"github.com/Oussamabh242/os-profile/bits"
	"github.com/Oussamabh242/os-profile/bits/keyboardevents"
	"github.com/Oussamabh242/os-profile/context"

	// "syscall/js"

	// "fmt"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/inpututil"
)

// var openURL = js.Global().Get("openURL")

var lastUpdate time.Time
var kr = keyboardevents.NewKeyRepeat()

const (
	SCREEN_WIDTH  = 1920
	SCREEN_HEIGHT = 1080
)

type Game struct {
	// visibleLines int
	ScreenText *bits.ScreenText
	ticks      int
	Controller uint8 // 0 -> tty , // 1 -> cat cmd
}

func (g *Game) Update() error {
	kr.Update()

	g.ticks++
	g.ScreenText.Update(g.ticks)

	// if kr.IsRepeat(ebiten.KeyControl) {
	// openURL.Invoke("https://example.com")
	// }

	keyboardevents.KeysGateway(g.ScreenText, kr, context.KeyToChar, &g.Controller)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if g.Controller == 1 {

		bits.DrawBase(screen, context.MD)
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

	//startmenu init

	// ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)

	ebiten.SetTPS(30)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
