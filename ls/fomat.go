package ls

import (
	"fmt"
	"my_ls_1/helper"
	"my_ls_1/parseArg"
)

// **************************** APPLIED THE FLAG****************
func GetTreeContent(path string) {
	helper.Totalsize=0
	helper.FilesInfo = Flags(path)
	if parseArg.Flags["a"] {
		helper.FilesInfo = Flagsa(path)
	}

	if parseArg.Flags["t"] {
		Flagst(path)
	}
	if parseArg.Flags["r"] {
		Flagsr()
	}
	Getfilename(helper.FilesInfo)

	if parseArg.Flags["l"] {
		Flagsl(path)
	}
}

func Resolve(path string) {
	is := false
	GetTreeContent(path)
	if parseArg.Flags["l"] {
		fmt.Println("Total", helper.Totalsize)
	}
	for _, v := range parseArg.Result {
		if parseArg.Flags["l"] {
			fmt.Println(v, " ")
		} else {
			is = true
			fmt.Print("\033[1m"+v+"\033[0m", " ")
		}
	}
	if is {
		fmt.Println()
	}
	parseArg.Result = []string{}

}
