--- 
Title = Rant on Typescript
Date = 21-08-2025 01:41
Author = Oussama Ben Hassen
---

# header one yes baby

yes i guess it is working

## Introduction 

in this Blog post we are going to make something worth our times, not just doing that 
one thing in your mind, no we are going beyond that 

### sub-sub-header

* some

~~~ts
    const x= 10
~~~


~~~go
func DrawMarkDown(md string, s *ebiten.Image) {
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
~~~ 

~~~rs
pub somego = 10;
~~~

~~~py

for i in range(10) : 
    print("whatever")
~~~

# some

