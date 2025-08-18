package mdparser

import (
	"regexp"
	"slices"
)

var StarterSignals []string = []string{"# ", "## ", "### ", "``` ", "* "}

func Parse(markdown string) []TextBlock {
	var blocks []TextBlock = []TextBlock{}

	// needed to start parsing
	mdLength := len(markdown)
	current := 0

	for current < mdLength {
		token, starterOffset := ParseNext(markdown, current, mdLength)
		blocks = append(blocks, token)
		current += len(token.Content) + starterOffset
	}
	return blocks
}

func ExistStarter(md string, cur int) string {

	starter := ""
	for i := cur + 2; i < cur+9 && i < len(md); i++ {
		// if i > len(md) {
		// 	return starter
		// }
		//
		possibleStarter := string(md[cur:i])

		xx := slices.Contains(StarterSignals, possibleStarter)
		if xx {
			starter = possibleStarter
			break
		}

		matched, _ := regexp.MatchString("^```[a-zA-Z0-9]{2}\n", possibleStarter)
		if matched == true {
			starter = possibleStarter
			break
		}

	}

	return starter
}

func ParseNext(markdown string, curr, length int) (TextBlock, int) {
	// minimum length of starters "# " || "* "
	starter := ExistStarter(markdown, curr)

	if starter == "# " {
		return TextBlock{
			Content: makeHeader(markdown, curr, 2),
			Type:    H1,
		}, 2
	}
	if starter == "## " {
		return TextBlock{
			Content: makeHeader(markdown, curr, 3),
			Type:    H2,
		}, 3
	}
	if starter == "### " {
		return TextBlock{
			Content: makeHeader(markdown, curr, 4),
			Type:    H3,
		}, 4
	}
	if starter == "* " {
		return TextBlock{
			Content: makeHeader(markdown, curr, 2),
			Type:    UL,
		}, 2
	}

	if match, _ := regexp.MatchString("^```[a-zA-Z0-9]{2}\n", starter); match == true {
		return TextBlock{
			Content: makeCodeBlock(markdown, curr, 6),
			Type:    CODE,
		}, 9

	}

	return TextBlock{
		Type:    NONE,
		Content: makeAny(markdown, curr),
	}, 0

}

func makeHeader(md string, cur, offset int) string {
	accumulated := []byte{}
	for i := cur + offset; i < len(md); i++ {
		if string(md[i]) == "\n" {
			break
		}
		accumulated = append(accumulated, md[i])
	}
	return string(accumulated)
}

func makeAny(md string, cur int) string {
	accumulated := []byte{}
	for i := cur; i < len(md); i++ {
		thisByte := md[i]
		accumulated = append(accumulated, thisByte)

		if string(thisByte) == "\n" && i+1 < len(md) {

			nextStarter := ""
			if i+1 < len(md) {
				nextStarter = ExistStarter(md, i+1)
			}

			if nextStarter != "" {
				break
			}

		}

	}
	return string(accumulated)
}

func makeCodeBlock(md string, cur, offset int) string {

	// lookup := md[cur+offset : cur+offset+3]
	accumulated := []byte{}
	for i := cur + offset; i < len(md); i++ {
		accumulated = append(accumulated, md[i])
		if i+3 < len(md) {
			if md[i:i+3] == "```" {
				break
			}
		}
	}
	return string(accumulated)

}
