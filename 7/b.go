package main

import (
	"fmt"
	"sort"

	"aoc-2022/7/common"
	"aoc-2022/pkg/utils"
)

const (
	totalDiskSpace   = 70000000
	availSpaceNeeded = 30000000
)

func main() {
	cmdLines := utils.ReadLines("7/input/real.txt")

	sepCmdLines := common.GetSeparatedCmdLines(cmdLines)

	cmds := utils.Map(sepCmdLines, common.NewCommand)

	// for _, cmd := range cmds {
	// 	fmt.Println(cmd.name, cmd.args, cmd.result)
	// }

	tree := common.ParseCommandsIntoTree(cmds)

	tree.ComputeSize()

	// fmt.Println()
	//fmt.Println(tree)
	files := tree.Filter(func(f common.FileTree) bool {
		return f.IsDir() && (totalDiskSpace-(tree.Size-f.Size) >= availSpaceNeeded)
	})

	fileSizes := utils.Map(files, func(f common.File) int {return int(f.Size)})
	sort.Ints(fileSizes)

	fmt.Println(fileSizes[0])

}
