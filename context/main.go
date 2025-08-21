package context

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"

	"github.com/Oussamabh242/os-profile/static"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type WorkDir struct {
	Path   string
	Parent string
}

var BASE_PROMPT = "oussama_ben_hassen@portfolio:"
var PROMPT = fmt.Sprintf("%s%s$ ", BASE_PROMPT, Wd.Path)
var Wd = WorkDir{
	Path:   "~",
	Parent: "~",
}

func UpdateDir(path string) {
	// PROMPT = fmt.Sprintf("%s%s $ ", BASE_PROMPT, path)

	PROMPT = fmt.Sprintf("%s%s $ ", BASE_PROMPT, path)
}

// func UpdateHead(node *Node) {
//
// }

const CURSOR = "_"

var BootText string = `
[    0.000000] Linux version 6.9.4-arch1-1 (linux@archlinux) (gcc (GCC) 13.2.1, GNU ld (GNU Binutils) 2.42) #1 SMP PREEMPT_DYNAMIC
[    0.000000] Command line: BOOT_IMAGE=/vmlinuz-linux root=/dev/sda2 rw quiet
[    2.118998] systemd[1]: Starting Load Kernel Modules...
[    2.344765] systemd[1]: Finished Load Kernel Modules.
[    2.518342] systemd[1]: Starting Login Service...
[    2.752004] systemd[1]: Started Login Service.
[    3.200390] systemd[1]: Startup finished in 2.014s (kernel) + 1.825s (userspace) = 3.839s.

Arch Linux  (tty1)

 ▄██████▄  ███    █▄     ▄████████    ▄████████    ▄████████   ▄▄▄▄███▄▄▄▄      ▄████████ 
███    ███ ███    ███   ███    ███   ███    ███   ███    ███ ▄██▀▀▀███▀▀▀██▄   ███    ███ 
███    ███ ███    ███   ███    █▀    ███    █▀    ███    ███ ███   ███   ███   ███    ███ 
███    ███ ███    ███   ███          ███          ███    ███ ███   ███   ███   ███    ███ 
███    ███ ███    ███ ▀███████████ ▀███████████ ▀███████████ ███   ███   ███ ▀███████████ 
███    ███ ███    ███          ███          ███   ███    ███ ███   ███   ███   ███    ███ 
███    ███ ███    ███    ▄█    ███    ▄█    ███   ███    ███ ███   ███   ███   ███    ███ 
 ▀██████▀  ████████▀   ▄████████▀   ▄████████▀    ███    █▀   ▀█   ███   █▀    ███    █▀  
                                                                                          
▀█████████▄     ▄████████ ███▄▄▄▄                                                         
  ███    ███   ███    ███ ███▀▀▀██▄                                                       
  ███    ███   ███    █▀  ███   ███                                                       
 ▄███▄▄▄██▀   ▄███▄▄▄     ███   ███                                                       
▀▀███▀▀▀██▄  ▀▀███▀▀▀     ███   ███                                                       
  ███    ██▄   ███    █▄  ███   ███                                                       
  ███    ███   ███    ███ ███   ███                                                       
▄█████████▀    ██████████  ▀█   █▀                                                        
                                                                                          
   ▄█    █▄       ▄████████    ▄████████    ▄████████    ▄████████ ███▄▄▄▄                
  ███    ███     ███    ███   ███    ███   ███    ███   ███    ███ ███▀▀▀██▄              
  ███    ███     ███    ███   ███    █▀    ███    █▀    ███    █▀  ███   ███              
 ▄███▄▄▄▄███▄▄   ███    ███   ███          ███         ▄███▄▄▄     ███   ███              
▀▀███▀▀▀▀███▀  ▀███████████ ▀███████████ ▀███████████ ▀▀███▀▀▀     ███   ███              
  ███    ███     ███    ███          ███          ███   ███    █▄  ███   ███              
  ███    ███     ███    ███    ▄█    ███    ▄█    ███   ███    ███ ███   ███              
  ███    █▀      ███    █▀   ▄████████▀   ▄████████▀    ██████████  ▀█   █▀               
                                                                                         


` + PROMPT

const Line_Y_OFFSET float64 = 25

var Head *Node = MakeFileTree()

const MAX_LINES = 35

const LS = ` 
drwxr-xr-x     - oussama_ben_hassen 25 Mar 16:10  CV 
drwxr-xr-x     - oussama_ben_hassen  4 Apr 12:09  Blog
-r--r--r--     - oussama_ben_hassen  4 Apr 12:09  Contact.md
`

var VisibleDirs map[string][]string = map[string][]string{
	"~": {"cv", "blog"},
}

var VisibleFiles map[string][]string = map[string][]string{
	"~":      {"Contact.md"},
	"~/blog": {"first.md", "second.md"},
}

var (
	FontIosevkaSource *text.GoTextFaceSource
)

var KeyToChar = map[ebiten.Key]string{
	ebiten.KeyA:     "a",
	ebiten.KeyB:     "b",
	ebiten.KeyC:     "c",
	ebiten.KeyD:     "d",
	ebiten.KeyE:     "e",
	ebiten.KeyF:     "f",
	ebiten.KeyG:     "g",
	ebiten.KeyH:     "h",
	ebiten.KeyI:     "i",
	ebiten.KeyJ:     "j",
	ebiten.KeyK:     "k",
	ebiten.KeyL:     "l",
	ebiten.KeyM:     "m",
	ebiten.KeyN:     "n",
	ebiten.KeyO:     "o",
	ebiten.KeyP:     "p",
	ebiten.KeyQ:     "q",
	ebiten.KeyR:     "r",
	ebiten.KeyS:     "s",
	ebiten.KeyT:     "t",
	ebiten.KeyU:     "u",
	ebiten.KeyV:     "v",
	ebiten.KeyW:     "w",
	ebiten.KeyX:     "x",
	ebiten.KeyY:     "y",
	ebiten.KeyZ:     "z",
	ebiten.KeySpace: " ",
	ebiten.KeyMinus: "-",
	// ebiten.KeyEnter: "\n",
	ebiten.Key0: "0",
	ebiten.Key1: "1",
	ebiten.Key2: "2",
	ebiten.Key3: "3",
	ebiten.Key4: "4",
	ebiten.Key5: "5",
	ebiten.Key6: "6",
	ebiten.Key7: "7",
	ebiten.Key8: "8",
	ebiten.Key9: "9",

	ebiten.KeyPeriod: ".",
}

func init() {

	Head.MakeNode("CV", Dir, nil)
	Blog := Head.MakeNode("Blog", Dir, nil)

	// SEARCH static/posts for blog posts
	files, _ := fs.ReadDir(static.PostsFS, "posts")
	// entries, err := os.ReadDir("static/posts")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {

	fmt.Println(files)
	for _, v := range files {

		if !v.IsDir() {

			filepath := "posts/" + v.Name()

			Blog.MakeNode(v.Name(), File, &filepath)
		}
	}

	// }

	Head.MakeNode("contact.md", File, nil)
	Blog.MakeNode("first.md", File, nil)
	Blog.MakeNode("second.md", File, nil)
	Blog.MakeNode("Some", Dir, nil)

	s, err := text.NewGoTextFaceSource(bytes.NewReader(static.IosevkaTTF))
	if err != nil {
		log.Fatal(err)
	}
	FontIosevkaSource = s

}

const NEOFETCH = `
	
                   -                    oussama_ben_hassen@portfolio
                  .o+                   -------------- 
                  ooo/                   OS: Arch Linux x86_64 
                 +oooo:                  Host: 486DX 
                +oooooo:                 Kernel: 2.4.36-arch1 
               -+oooooo+:                Uptime: 5 days, 7 hours 
              /:-:++oooo+:               Packages: 387 (pacman) 
             /++++/+++++++:              Shell: bash 2.05 
            /++++++++++++++:             Terminal: tty1 
           /+++ooooooooooooo/            DE: None 
         ./ooosssso++osssssso+           WM: None 
        .oossssso-    /ossssss+          Theme: None 
       -osssssso.      :ssssssso.        Icons: None 
      :osssssss/        osssso+++.       CPU: Intel 486DX @ 66MHz 
     /ossssssss/        +ssssooo/-       GPU: VGA compatible 
    /ossssso+/:-        -:/+osssso+-     Memory: 31MiB / 64MiB 
   +sso+:-                   .-/+oso:    Swap: 64MiB / 128MiB 
  ++:.                            -/+/
 .                                   /                           

	`

var MD string = `# header one yes baby

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
`

var CatStartY = 10

var KEYWORDS []string = []string{
	// Go Keywords
	"break", "case", "chan", "const", "continue",
	"default", "defer", "else", "fallthrough", "for",
	"func", "go", "goto", "if", "import",
	"interface", "map", "package", "range", "return",
	"select", "struct", "switch", "type", "var",

	// Lua Keywords
	"and", "do", "elseif", "end", "false",
	"for", "function", "goto", "if", "in",
	"local", "nil", "not", "or", "repeat",
	"return", "then", "true", "until", "while",
	"break", "else",

	// TypeScript Keywords
	"any", "as", "async", "await", "boolean",
	"catch", "class", "constructor", "declare", "default",
	"delete", "enum", "export", "extends", "false",
	"finally", "from", "function", "get", "if",
	"implements", "import", "in", "instanceof", "interface",
	"keyof", "let", "module", "namespace", "never",
	"new", "null", "number", "private", "protected",
	"public", "readonly", "require", "return", "set",
	"static", "string", "super", "switch", "symbol",
	"this", "throw", "true", "try", "type",
	"typeof", "undefined", "unknown", "var", "void",
	"while", "with", "yield",

	// Bash Keywords
	"!", "do", "done", "elif", "else",
	"esac", "fi", "for", "function", "if",
	"in", "select", "then", "time", "until",
	"while",

	// Coreutils & Common CLI Tools
	"ls", "cp", "mv", "rm", "mkdir",
	"rmdir", "touch", "stat", "file", "find",
	"basename", "dirname",

	"cat", "grep", "sed", "awk", "cut",
	"sort", "uniq", "tr", "head", "tail",
	"wc", "xargs", "split", "paste", "tee",

	"chmod", "chown", "chgrp", "umask",

	"top", "ps", "df", "du", "free",
	"uptime", "vmstat", "iostat", "who", "w",

	// CLI Tools
	"curl", "wget", "ssh", "scp", "rsync",
	"tar", "zip", "unzip", "gzip", "bzip2",
	"jq", "tree", "watch", "less", "more",
	"htop", "ping", "traceroute",

	"git", "docker", "kubectl", "nmap", "make",
	"npm", "yarn", "node", "go", "ts-node", "deno",
	"lua", "cargo", "rustc", "pip", "python",

	// Bonus Bash Tools
	"fzf", "bat", "exa", "fd", "rg",
	"tldr", "ncdu", "btop", "lsd", "httpie",
}
