package context

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFileTree(t *testing.T) {
	ft := MakeFileTree()
	ft.MakeNode("About", Dir)
	ft.MakeNode("temp.txt", File)
	ft.MakeNode("temp.txt", File)
	fmt.Println(strings.Join(ft.Ls(), "\n"))
	if ft.Parent != nil {
		t.Errorf("root's parent is not nil ")
	}
	pwd := Pwd(ft)
	fmt.Println("pwd : ", pwd)
	assert.Equal(t, "~", pwd, "expected '/' and got "+pwd)

	// chdir
	About, err := ft.Cd("About")
	assert.Equal(t, "", err)
	_, err = ft.Cd("some")
	assert.Greater(t, len(err), 0)

	pwd = Pwd(About)
	fmt.Println(pwd)
	assert.Equal(t, "~/About", pwd)

}
