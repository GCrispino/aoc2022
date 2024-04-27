package common

import (
	"fmt"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size uint
}

func (f *File) String() string {
	return fmt.Sprintf("%s - %d", f.Name, f.Size)
}

type FileTree struct {
	*File
	children []*FileTree
    parent *FileTree
}

func (f FileTree) IsDir() bool {
	return len(f.children) > 0
}

func (f FileTree) String() string {
	s := fmt.Sprintf("%s - %d. [", f.Name, f.Size)

	for i, child := range f.children {
		s += fmt.Sprintf("%s - %d", child.Name, child.Size)
		if child.IsDir() {
			for _, _child := range child.children {
				s += " [" + _child.String() + "] "
			}
		}
		if i < len(f.children)-1 {
			s += ","
		}
	}
	s += "]"
	return s
}

func (f FileTree) ComputeSize() {
	Size := uint(0)
	for _, child := range f.children {
		if child.IsDir() {
			child.ComputeSize()
		}

		Size += child.Size
	}

	f.Size = Size
}

func (f FileTree) Filter(fn func (FileTree) bool) []File {
	files := make([]File, 0)

	if fn(f) {
		files = append(files, *f.File)
	}

	for _, child := range f.children {
		files = append(files, child.Filter(fn)...)
	}

	return files
}

type command struct {
	Name   string
	args   []string
	result []string
}

func NewCommand(cmdStrs []string) command {
	originalCmd := cmdStrs[0]
	cmdResult := cmdStrs[1:]

	spl := strings.Split(originalCmd[2:], " ")

	return command{
		Name:   spl[0],
		args:   spl[1:],
		result: cmdResult,
	}
}

const (
	cdCommand = "cd"
	lsCommand = "ls"
)

// type cdCommand struct {
// 	command
// }
//
// type lsCommand struct {
// 	command
// }

func ParseCommandsIntoTree(commands []command) FileTree {
	tree := FileTree{
		children: make([]*FileTree, 0),
	}
	treePtr := &tree
    
	for _, command := range commands {
		// fmt.Println("  treePtr:", treePtr)
		// fmt.Printf("    %s command with args %v\n", command.Name, command.args)
		switch command.Name {
		case cdCommand:
			arg := command.args[0]
            if arg == ".." {
              treePtr = treePtr.parent
              break
            }
			if tree.File == nil {
				tree.File = &File{Name: arg}
			} else {
				for _, child := range treePtr.children {
                    // fmt.Println("      child:", child)
					if child.File.Name == arg {
						// fmt.Println("  achou filho. Nome: ", child.Name)
						treePtr = child
                        break
					}
				}
			}
		case lsCommand:
			for _, fileListing := range command.result {
				spl := strings.Split(fileListing, " ")

				f := FileTree{
                  children: make([]*FileTree, 0),
                  parent: treePtr,
                }
				if spl[0] == "dir" {
					f.File = &File{Name: spl[1]}
				} else { // ls
					Size, err := strconv.ParseUint(spl[0], 10, 64)
					if err != nil {
						panic(err)
					}
					f.File = &File{
						Name: spl[1],
						Size: uint(Size),
					}
				}
				treePtr.children = append(treePtr.children, &f)
                // fmt.Println("  new children", treePtr.children)
			}
		default:
			panic("invalid command Name")
		}
	}

	return tree
}

func GetSeparatedCmdLines(cmdLines []string) [][]string {
	sepCmdLines := make([][]string, 0)

	var curCmd []string
	for i, line := range cmdLines {
		if line[0] == '$' {
          // fmt.Println(line)
			if len(curCmd) > 0 {
				sepCmdLines = append(sepCmdLines, curCmd)
			}
			curCmd = []string{line}
		} else {
			curCmd = append(curCmd, line)
		}

        if i == len(cmdLines) - 1 {
				sepCmdLines = append(sepCmdLines, curCmd)
        }
	}

	return sepCmdLines
}
