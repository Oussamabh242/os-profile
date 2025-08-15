package bits

import (
	"exmp/context"
	"fmt"
	"image/color"
	"log"
	"slices"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ScreenText struct {
	Text         string
	TextLength   int
	booted       bool
	visibleLines int
}

func NewScreenText() *ScreenText {
	return &ScreenText{
		Text:         context.BootText,
		TextLength:   len(context.BootText),
		booted:       false,
		visibleLines: 0,
	}
}

func (st *ScreenText) Update(ticks int) {
	if ticks%1 == 0 {
		if st.visibleLines < len(strings.Split(context.BootText, "\n")) {
			st.visibleLines++
		}
	}
	if st.visibleLines == len(strings.Split(context.BootText, "\n")) {
		st.booted = true
		st.visibleLines = 41
	}
	// cursor blink
	if st.booted {

		if ticks%20 == 0 {

			lastchar := st.Text[len(st.Text)-1]
			textLength := len(st.Text)

			if string(lastchar) == context.CURSOR {
				st.Text = st.Text[0 : textLength-1]
				st.TextLength--
			} else {
				st.Text = st.Text + context.CURSOR
				st.TextLength++
			}

		}

	}
}

func (st *ScreenText) Draw(screen *ebiten.Image) {
	// if st.booted && len(st.Text) == st.TextLength {
	// 	return
	// }
	lines := strings.Split(st.Text, "\n")
	for i := 0; i < st.visibleLines && i < len(lines); i++ {
		line := lines[i]
		op := &text.DrawOptions{}
		op.GeoM.Translate(20, float64(i)*context.Line_Y_OFFSET)
		op.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, line, &text.GoTextFace{
			Source: context.FontIosevkaSource,
			Size:   24,
		}, op)
	}

}

func (st *ScreenText) AddChar(c string) {
	lastchar := st.Text[len(st.Text)-1]
	textLength := len(st.Text)

	lines := strings.Split(st.Text, "\n")
	newLine := ""
	if len(lines[len(lines)-1]) > 100 {
		newLine = "\n"
		fmt.Println("yes")
	}

	if string(lastchar) == context.CURSOR {
		st.Text = st.Text[0:textLength-1] + newLine + c + context.CURSOR
	} else {
		st.Text = st.Text + newLine + c
	}
	if newLine == "\n" {
		fmt.Println(st.Text)
	}
}

func (st *ScreenText) DeleteChar() {
	lastchar := st.Text[len(st.Text)-1]
	textLength := len(st.Text)

	lines := strings.Split(st.Text, "\n")
	lastLine := lines[len(lines)-1]
	if string(lastLine[len(lastLine)-1]) == context.CURSOR {
		lastLine = lastLine[0 : len(lastLine)-1]
	}

	if lastLine == context.PROMPT {
		return
	}

	if string(lastchar) == context.CURSOR {
		st.Text = st.Text[0:textLength-2] + context.CURSOR
	} else {
		st.Text = st.Text[0 : textLength-1]
	}
}

func (st *ScreenText) Execute() {

	lines := strings.Split(st.Text, "\n")
	lastLine := lines[len(lines)-1]
	lastchar := lastLine[len(lastLine)-1]

	if string(lastchar) == context.CURSOR {
		lastLine = lastLine[:len(lastLine)-1]
		st.Text = st.Text[:len(st.Text)-1]
	}

	cmd := lastLine[len(context.PROMPT):]
	nt := commandsGateway(st, cmd, lines)

	st.Text = nt + "\n" + context.PROMPT + context.CURSOR
	st.cleanUp()
	// st.visibleLines++
}
func (st *ScreenText) cleanUp() {
	lines := strings.Split(st.Text, "\n")
	if len(lines) < 40 {
		return
	}

	st.Text = strings.Join(lines[len(lines)-40:len(lines)], "\n")

}

func commandsGateway(st *ScreenText, cmd string, lines []string) string {

	if len(cmd) > 2 && cmd[0:3] == "cd " {
		newdir := cmd[3:]
		fmt.Println(newdir, len(newdir), newdir == "..")
		if newdir == ".." || newdir == "../" {

			context.Wd.Path = "~"
			context.Wd.Parent = "~"
			context.UpdateDir(context.Wd.Path)
			return st.Text

		}
		if !slices.Contains(context.VisibleDirs[context.Wd.Path], newdir) {
			return st.Text + "\n" + "cd: no such file or directory: " + newdir
		}
		context.Wd.Parent = context.Wd.Path
		context.Wd.Path += "/" + newdir

		context.UpdateDir(context.Wd.Path)
		return st.Text
	}

	if cmd == "clear" {
		return ""
	}

	if cmd == "ls" {
		vf := context.VisibleFiles[context.Wd.Path]
		vd := context.VisibleDirs[context.Wd.Path]
		combined := ""
		for _, val := range vd {
			combined += "drwxr-xr-x     - oussama_ben_hassen   " + val + "\n"
		}
		for _, val := range vf {
			combined += "-r--r--r--		  - oussama_ben_hassen   " + val + "\n"
		}
		return st.Text + "\n" + combined
	}
	if cmd == "pwd" {
		return st.Text + "\n" + context.Wd.Path
	}

	if cmd == "neofetch" {
		return st.Text + context.NEOFETCH
	}

	if cmd == "" {
		return st.Text
	}
	if cmd == "exit" {
		log.Fatal()
	}
	return st.Text + "\n" + "bash: " + cmd + " : command not found"

}

func CmdChangeDir(cmd string, stText string) {

}
