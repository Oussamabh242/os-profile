package bits

import (
	"image/color"
	"slices"
	"strings"

	"github.com/Oussamabh242/os-profile/context"
	mdparser "github.com/Oussamabh242/os-profile/md-parser"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func DrawBase(s *ebiten.Image, md string) {

	x := ebiten.NewImage(1920, 30)
	x.Fill(color.RGBA{255, 0, 255, 255})

	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(0, 2) // Adjust Y so it's within 30px height
	textOp.ColorScale.ScaleWithColor(color.White)

	text.Draw(x, "            q : quit            | n : scroll down            | p : scroll up", &text.GoTextFace{
		Source: context.FontIosevkaSource,
		Size:   20,
	}, textOp)

	DrawMarkDown(md, s)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 1040)
	s.DrawImage(x, op)

	// s.Fill(color.Black)
	// x := ebiten.NewImage(1920, 30)
	// x.Fill(color.RGBA{255, 0, 255, 255})
	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(0, 1040)
	//
	// textOp := &text.DrawOptions{}
	// textOp.GeoM.Translate(0, 1042)
	// textOp.ColorScale.ScaleWithColor(color.White)
	//
	// text.Draw(x, "            q : quit            | n : scroll down            | p : scroll up", &text.GoTextFace{
	// 	Source: context.FontIosevkaSource,
	// 	Size:   20,
	// }, textOp)
	//
	// s.DrawImage(x, op)
}

func DrawMarkDown(md string, s *ebiten.Image) {
	startx := 5
	starty := context.CatStartY
	parts := mdparser.Parse(md)
	for _, v := range parts {
		lineSize := typeToSize(v.Type)
		splitted := strings.Split(v.Content, "\n")
		for _, xline := range splitted {
			line := v.LangToLogo() + "\n" + xline
			backlineCount := strings.Count(line, "\n")
			if backlineCount == 0 {
				backlineCount = 1
			}
			box := ebiten.NewImage(1920, (lineSize+7)*backlineCount)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(startx), float64(starty))

			// Process line while preserving whitespace
			xOffset := 0
			charWidth := int(float64(lineSize) * 0.6)

			// Find keyword positions for coloring
			keywordPositions := make(map[int]bool)

			// Check each keyword in the KEYWORDS slice
			for _, keyword := range context.KEYWORDS { // Assuming KEYWORDS is accessible via context
				keywordIndex := 0
				for {
					pos := strings.Index(line[keywordIndex:], keyword)
					if pos == -1 {
						break
					}
					realPos := keywordIndex + pos

					// Check if it's a whole word (not part of another word)
					isWholeWord := true
					if realPos > 0 {
						prevChar := line[realPos-1]
						if (prevChar >= 'a' && prevChar <= 'z') ||
							(prevChar >= 'A' && prevChar <= 'Z') ||
							(prevChar >= '0' && prevChar <= '9') ||
							prevChar == '_' {
							isWholeWord = false
						}
					}
					if realPos+len(keyword) < len(line) {
						nextChar := line[realPos+len(keyword)]
						if (nextChar >= 'a' && nextChar <= 'z') ||
							(nextChar >= 'A' && nextChar <= 'Z') ||
							(nextChar >= '0' && nextChar <= '9') ||
							nextChar == '_' {
							isWholeWord = false
						}
					}

					if isWholeWord {
						for i := realPos; i < realPos+len(keyword) && i < len(line); i++ {
							keywordPositions[i] = true
						}
					}
					keywordIndex = realPos + 1
				}
			}

			// Draw character by character
			for i, char := range line {
				if char == '\t' {
					// Handle tab as 4 spaces
					xOffset += charWidth * 4
					continue
				}

				charStr := string(char)
				textOp := &text.DrawOptions{}
				textOp.GeoM.Translate(float64(xOffset), 2)

				if keywordPositions[i] && v.Type == mdparser.CODE {
					textOp.ColorScale.ScaleWithColor(color.RGBA{100, 149, 237, 255})
				} else {
					textOp.ColorScale.ScaleWithColor(color.White)
				}

				text.Draw(box, charStr, &text.GoTextFace{
					Source: context.FontIosevkaSource,
					Size:   float64(lineSize),
				}, textOp)

				xOffset += charWidth
			}

			s.DrawImage(box, op)
			starty += (lineSize + 7) * backlineCount
		}
	}
}

func DrawMarkDown4(md string, s *ebiten.Image) {
	startx := 5
	starty := context.CatStartY
	parts := mdparser.Parse(md)
	for _, v := range parts {
		lineSize := typeToSize(v.Type)
		splitted := strings.Split(v.Content, "\n")
		for _, line := range splitted {
			backlineCount := strings.Count(line, "\n")
			if backlineCount == 0 {
				backlineCount = 1
			}
			box := ebiten.NewImage(1920, (lineSize+7)*backlineCount)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(startx), float64(starty))

			// Process line while preserving whitespace
			xOffset := 0
			charWidth := int(float64(lineSize) * 0.6)

			// Find ERROR positions for coloring
			errorPositions := make(map[int]bool)
			errorIndex := 0
			for {
				pos := strings.Index(line[errorIndex:], "ERROR")
				if pos == -1 {
					break
				}
				realPos := errorIndex + pos
				for i := realPos; i < realPos+5 && i < len(line); i++ {
					errorPositions[i] = true
				}
				errorIndex = realPos + 1
			}

			// Draw character by character
			for i, char := range line {
				if char == '\t' {
					// Handle tab as 4 spaces
					xOffset += charWidth * 4
					continue
				}

				charStr := string(char)
				textOp := &text.DrawOptions{}
				textOp.GeoM.Translate(float64(xOffset), 2)

				if errorPositions[i] {
					textOp.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
				} else {
					textOp.ColorScale.ScaleWithColor(color.White)
				}

				text.Draw(box, charStr, &text.GoTextFace{
					Source: context.FontIosevkaSource,
					Size:   float64(lineSize),
				}, textOp)

				xOffset += charWidth
			}

			s.DrawImage(box, op)
			starty += (lineSize + 7) * backlineCount
		}
	}
}

// op := &text.DrawOptions{}
//
//	op.GeoM.Translate(20, float64(i)*context.Line_Y_OFFSET)
//	op.ColorScale.ScaleWithColor(color.White)
//	text.Draw(screen, line, &text.GoTextFace{
//		Source: context.FontIosevkaSource,
//		Size:   24,
//	}, op)
func DrawMarkDown3(md string, s *ebiten.Image) {
	startx := 5
	starty := context.CatStartY
	parts := mdparser.Parse(md)
	for _, v := range parts {
		lineSize := typeToSize(v.Type)
		splitted := strings.Split(v.Content, "\n")
		for _, line := range splitted {
			backlineCount := strings.Count(line, "\n")
			if backlineCount == 0 {
				backlineCount = 1
			}
			box := ebiten.NewImage(1920, (lineSize+7)*backlineCount)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(startx), float64(starty))

			// Only draw word by word (remove the duplicate line drawing)
			words := strings.Fields(line)
			xOffset := 0
			for _, word := range words {
				wordWithSpace := word + " "
				textOp := &text.DrawOptions{}
				textOp.GeoM.Translate(float64(xOffset), 2)

				// fmt.Printf("%q\n", word)
				if v.Type == mdparser.CODE && slices.Contains(context.KEYWORDS, word) {
					textOp.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
				} else {
					textOp.ColorScale.ScaleWithColor(color.White)
				}
				text.Draw(box, wordWithSpace, &text.GoTextFace{
					Source: context.FontIosevkaSource,
					Size:   float64(lineSize),
				}, textOp)

				charWidth := int(float64(lineSize) * 0.6)
				approxAdvance := len(wordWithSpace) * charWidth
				xOffset += approxAdvance
			}

			s.DrawImage(box, op)
			starty += (lineSize + 7) * backlineCount
		}
	}
}

func DrawMarkDown2(md string, s *ebiten.Image) {
	startx := 5
	starty := context.CatStartY
	parts := mdparser.Parse(md)
	for _, v := range parts {

		// if v.Type == mdparser.CODE {
		//
		// 	v.Content += "\n\n"
		// 	fmt.Printf("%q\n", v.Content)
		// }

		lineSize := typeToSize(v.Type)
		splitted := strings.Split(v.Content, "\n")

		for _, line := range splitted {

			backlineCount := strings.Count(line, "\n")
			if backlineCount == 0 {
				backlineCount = 1
			}

			box := ebiten.NewImage(1920, (lineSize+7)*backlineCount) // AKA allocating space to write text
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(startx), float64(starty))

			words := strings.Fields(line)
			xOffset := 0

			for _, word := range words {
				wordWithSpace := word + " "

				textOp := &text.DrawOptions{}
				textOp.GeoM.Translate(float64(xOffset), 2)

				if word == "const" {
					textOp.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
				} else {
					textOp.ColorScale.ScaleWithColor(color.White)
				}
				text.Draw(box, wordWithSpace, &text.GoTextFace{
					Source: context.FontIosevkaSource,
					Size:   float64(lineSize),
				}, textOp)

				// face := context.FontIosevkaSource.Face

				charWidth := int(float64(lineSize) * 0.6) // tweak 0.55 to 0.65 based on visual testing
				approxAdvance := len(wordWithSpace) * charWidth

				xOffset += approxAdvance
			}

			textOp := &text.DrawOptions{}
			textOp.GeoM.Translate(0, 2)
			text.Draw(box, line, &text.GoTextFace{
				Source: context.FontIosevkaSource,
				Size:   float64(lineSize),
			}, textOp)

			s.DrawImage(box, op)

			starty += (lineSize + 7) * backlineCount

		}
	}
}

func typeToSize(t mdparser.BlockType) int {
	switch t {
	case mdparser.H1:
		return 45

	case mdparser.H2:
		return 40
	case mdparser.H3:
		return 35
	case mdparser.NONE:
		return 24
	case mdparser.CODE:
		return 30
	default:
		return 24
	}
}
