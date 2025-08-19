package mdparser

type Renderable interface{}

type BlockType int

const (
	H1 BlockType = iota
	H2
	H3
	UL
	NONE
	CODE
	IMG
)

// TextBlock can be a single line of text eg Header
// or an unorder list * ..
// or even a block of code.
type TextBlock struct {
	Type    BlockType
	Content string
	Lang    string
}

func NewText(xtype BlockType, content string) TextBlock {
	return TextBlock{
		Type:    xtype,
		Content: content,
	}
}

var Starters []string = []string{"# ", "## ", "### ", "``` ", "* "}
