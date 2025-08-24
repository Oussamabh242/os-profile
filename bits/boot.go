package bits

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"strings"

	"github.com/Oussamabh242/os-profile/context"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type ScreenText struct {
	Text         string
	ArrayText    []string
	TextLength   int
	booted       bool
	visibleLines int
}

func NewScreenText() *ScreenText {
	return &ScreenText{
		Text:         context.BootText,
		ArrayText:    strings.Split(context.BootText, "\n"),
		TextLength:   len(context.BootText),
		booted:       false,
		visibleLines: 0,
	}
}
func (st *ScreenText) RecomputeArray() {
	st.ArrayText = strings.Split(st.Text, "\n")
}

func (st *ScreenText) Update(ticks int) {
	splittedText := strings.Split(context.BootText, "\n")
	if ticks%1 == 0 {
		if st.visibleLines < len(splittedText) {
			st.visibleLines++
		}
	}
	if st.visibleLines == len(splittedText) {
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

			st.RecomputeArray()
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

	lines := st.ArrayText
	newLine := ""
	if len(lines[len(lines)-1]) > 100 {
		newLine = "\n"
	}

	if string(lastchar) == context.CURSOR {
		st.Text = st.Text[0:textLength-1] + newLine + c + context.CURSOR
	} else {
		st.Text = st.Text + newLine + c
	}
	if newLine == "\n" {
	}
	st.RecomputeArray()
}

func (st *ScreenText) DeleteChar() {
	lastchar := st.Text[len(st.Text)-1]
	textLength := len(st.Text)

	lines := st.ArrayText
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
	st.RecomputeArray()
}

func (st *ScreenText) Execute(controller *uint8, fileHoldingControl *string) {

	lines := st.ArrayText
	lastLine := lines[len(lines)-1]
	lastchar := lastLine[len(lastLine)-1]

	if string(lastchar) == context.CURSOR {
		lastLine = lastLine[:len(lastLine)-1]
		st.Text = st.Text[:len(st.Text)-1]
	}

	cmd := lastLine[len(context.PROMPT):]
	fmt.Println(cmd)
	nt := commandsGateway(st, cmd, lines, controller, fileHoldingControl)

	st.Text = nt + "\n" + context.PROMPT + context.CURSOR
	st.cleanUp()
	// st.visibleLines++

	st.RecomputeArray()
}
func (st *ScreenText) cleanUp() {
	lines := strings.Split(st.Text, "\n")
	if len(lines) < 40 {
		return
	}

	st.Text = strings.Join(lines[len(lines)-40:len(lines)], "\n")
	st.RecomputeArray()

}

func commandsGateway(st *ScreenText, cmd string, lines []string, controller *uint8,
	fileHoldingControl *string) string {

	splitted := strings.Split(cmd, " ")

	if len(splitted) == 2 && splitted[0] == "srch" && context.Head.Name == "Blog" {
		contains := splitted[1]
		str := context.MakeOutsideReqeust("/posts?search=" + contains)

		var ret context.ListPostsReq
		json.Unmarshal([]byte(str), &ret)

		if len(ret.Posts) == 0 {
			return st.Text + "\n" + "srch: " + contains + " : no file matching this pattern found! "
		}

		matchs := ""
		mxcols := 5

		for _, v := range ret.Posts {
			if mxcols == 0 {
				mxcols = 5
				matchs += "\n"
			}
			matchs += v + "        "
			mxcols--
		}
		return st.Text + "\n" + matchs

	}

	if len(splitted) == 2 && splitted[0] == "cat" {
		fileName := splitted[1]
		// filesInCurrentDir := context.Head.Ls()
		exists := context.Node{}
		for _, v := range context.Head.Children {
			if v.Type == context.File && v.Name == fileName {
				exists = *v
				break
			}
		}
		if exists.Name == "" {
			return st.Text + "\n" + "File : " + fileName + " does not exists in the current directory" + "\n"
		}
		*controller = 1
		*fileHoldingControl = exists.Name
		return st.Text
	}

	if len(cmd) > 2 && cmd[0:3] == "cd " {
		newdir := cmd[3:]

		var errs string
		newHead, errs := context.Head.Cd(newdir)
		if errs != "" {
		} else {
			context.Head = newHead
			context.UpdateDir(context.Pwd(context.Head))

		}

		// if newdir == ".." || newdir == "../" {
		//
		// 	context.Wd.Path = "~"
		// 	context.Wd.Parent = "~"

		//
		// 	return st.Text
		//
		// }
		// if !slices.Contains(context.VisibleDirs[context.Wd.Path], newdir) {
		// 	return st.Text + "\n" + "cd: no such file or directory: " + newdir
		// }
		// context.Wd.Parent = context.Wd.Path
		// context.Wd.Path += "/" + newdir
		//
		// context.UpdateDir(context.Wd.Path)
		return st.Text
	}

	if cmd == "clear" {
		return ""
	}

	if cmd == "ls" {
		return st.Text + "\n" + strings.Join(context.Head.Ls(), "\n")
		// vf := context.VisibleFiles[context.Wd.Path]
		// vd := context.VisibleDirs[context.Wd.Path]
		// combined := ""
		// for _, val := range vd {
		// 	combined += "drwxr-xr-x     - oussama_ben_hassen   " + val + "\n"
		// }
		// for _, val := range vf {
		// 	combined += "-r--r--r--		  - oussama_ben_hassen   " + val + "\n"
		// }
		// return st.Text + "\n" + combined
	}
	if cmd == "pwd" {
		return st.Text + "\n" + context.Pwd(context.Head)
		// return st.Text + "\n" + context.Wd.Path
	}

	if cmd == "neofetch" {
		return st.Text + context.NEOFETCH
	}

	if cmd == "" {
		return st.Text
	}

	// if cmd == "cat" {
	// 	*controller = 1
	// 	return st.Text
	// }

	if cmd == "exit" {
		log.Fatal()
	}
	return st.Text + "\n" + "bash: " + cmd + " : command not found"

}
