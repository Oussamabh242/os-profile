package context

import (
	"bytes"
	_ "embed"
	"exmp/static"
	"fmt"
	"log"

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
[    0.000000] x86/fpu: Supporting XSAVE feature 0x001: 'x87 floating point registers'
[    0.000000] x86/fpu: Enabled xstate features 0x1, context size is 0x240 bytes, using 'compacted' format.
[    0.000000] BIOS-e820: [mem 0x0000000000000000-0x000000000009ffff] usable
...
[    1.245620] systemd[1]: Detected architecture x86-64.
[    1.248012] systemd[1]: Hostname set to <archlinux>.
[    1.455789] systemd[1]: Reached target Local File Systems.
[    1.478120] systemd[1]: Starting Journal Service...
[    1.592130] systemd-journald[234]: Journal started
[    2.005200] systemd[1]: Started Journal Service.
[    2.118998] systemd[1]: Starting Load Kernel Modules...
[    2.344765] systemd[1]: Finished Load Kernel Modules.
[    2.518342] systemd[1]: Starting Login Service...
[    2.712034] systemd-logind[376]: New seat seat0.
[    2.752004] systemd[1]: Started Login Service.
[    3.000123] systemd[1]: Reached target Multi-User System.
[    3.143003] systemd[1]: Reached target Graphical Interface.
[    3.200390] systemd[1]: Startup finished in 2.014s (kernel) + 1.825s (userspace) = 3.839s.

Arch Linux  (tty1)


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

	Head.MakeNode("CV", Dir)
	Blog := Head.MakeNode("Blog", Dir)
	Head.MakeNode("contact.md", File)
	Blog.MakeNode("first.md", File)
	Blog.MakeNode("second.md", File)
	Blog.MakeNode("Some", Dir)

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
