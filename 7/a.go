package day7

import (
	"fmt"

	"aoc-2022/7/common"
	"aoc-2022/pkg/utils"
)

func SolveA() {
	cmdLines := utils.ReadLines("7/input/real.txt")

	sepCmdLines := common.GetSeparatedCmdLines(cmdLines)

	cmds := utils.Map(sepCmdLines, common.NewCommand)

	// for _, cmd := range cmds {
	// 	fmt.Println(cmd.name, cmd.args, cmd.result)
	// }

	tree := common.ParseCommandsIntoTree(cmds)

	tree.ComputeSize()

	// fmt.Println()
	// fmt.Println(tree)
	files := tree.Filter(func(f common.FileTree) bool {
		return f.IsDir() && f.Size < 100000
	})

	// fmt.Println("files:", files)

	sum := uint(0)
	for _, file := range files {
		sum += file.Size
	}
	fmt.Println(sum)
	// for _, child := range tree.children {
	//   fmt.Println("  ", child)
	// }
}
