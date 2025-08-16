package context

import "fmt"

type Filetype int

const (
	File Filetype = iota
	Dir
)

type Node struct {
	Name     string
	Type     Filetype
	Children []*Node
	IsRoot   bool
	Parent   *Node
}

// type WhereAmI struct {
// 	CWD *Node
// }

func MakeFileTree() *Node {
	return &Node{
		Name:     "~",
		Type:     Dir,
		Children: []*Node{},
		IsRoot:   true,
	}
}

func (n *Node) MakeNode(name string, nodeType Filetype) *Node {

	newNode := &Node{
		Name:     name,
		Type:     nodeType,
		Children: []*Node{},
		IsRoot:   false,
		Parent:   n,
	}

	n.Children = append(n.Children, newNode)
	return newNode

}

func containsDirName(nodes []*Node, nodeName string) *Node {
	for _, v := range nodes {
		if v.Name == nodeName && v.Type == Dir {
			return v
		}
	}
	return nil
}

func (n *Node) Ls() []string {

	curlist := []string{}

	for _, node := range n.Children {
		//-r--r--r--     - oussama_ben_hassen  4 Apr 12:09

		if node.Type == File {
			curlist = append(curlist, "-r--r--r--     - oussama_ben_hassen  30 Apr 12:09\t "+node.Name)
		} else {
			curlist = append(curlist, "dr-xr-xr-x     - oussama_ben_hassen  30 Apr 12:09\t "+node.Name)
		}
	}
	return curlist
}

func (n *Node) Cd(dir string) (*Node, string) {
	fmt.Println("inside CD-FUNC : " + dir)
	if dir == ".." && n.IsRoot == false && n.Parent != nil {
		return n.Parent, ""
	}

	fmt.Println("inside CD-FUNC : " + dir)

	for _, node := range n.Children {
		//-r--r--r--     - oussama_ben_hassen  4 Apr 12:09
		fmt.Println("inside CD-FUNC : ", node.Name)

		if node.Type == File {
			continue
		}

		if node.Name == dir && node.Type == Dir {
			return node, ""
		}
		if node.Name == dir && node.Type != Dir {

			return nil, "cd: not a directory: " + dir
		}

	}
	return nil, "cd: no such file or directory: " + dir
}

func Pwd(node *Node) string {
	temp := node

	pwd := ""
	if node.IsRoot == true {
		return "~"
	}

	for temp != nil {
		if temp.IsRoot == true {
			pwd = "~" + pwd
			break
		}
		pwd = "/" + temp.Name + pwd
		temp = temp.Parent
	}

	return pwd

}
