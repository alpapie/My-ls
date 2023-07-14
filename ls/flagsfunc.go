package ls

import (
	"io/fs"
	"my_ls_1/helper"
	"my_ls_1/parseArg"
	"os"
)

var Colors = map[string]string{
	"blue":  "\033[38;2;0;122;204m",
	"reset": "\033[0m",
}

// ***************************** THE -A(list the hadden folder) FLAG *************************
func Flagsa(_path string) []helper.FilesEntry {
	parseArg.Result = []string{}
	files := []helper.FilesEntry{}
	for i, v := range []string{_path, _path + "/.."} {
		// point, _ := os.DirFS(v).Open(".")
		point, _ := os.DirFS(v).Open(".")
		_current, _ := point.Stat()
		current := fs.FileInfoToDirEntry(_current)
		filehadden := helper.FilesEntry{EntryFile: current}
		if i == 1 {
			filehadden.P_name = ".."
		} else {
			filehadden.P_name = "."
		}
		files = append(files, filehadden)
	}

	files = append(files, helper.GetFile(_path, true)...)
	return files
}

// ***************************** THE -t(sort by time) FLAG *************************
func Flagst(_path string) {
	length := len(helper.FilesInfo)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			a, _ := helper.FilesInfo[i].EntryFile.Info()
			b, _ := helper.FilesInfo[j].EntryFile.Info()
			// if i==0{
			// 	fmt.Println(b.ModTime())
			// }
			if a.ModTime().Before(b.ModTime()) {
				helper.FilesInfo[i], helper.FilesInfo[j] = helper.FilesInfo[j], helper.FilesInfo[i]
			}
		}
	}
}

// ***************************** THE -r(reverse the ouput) FLAG *************************
func Flagsr() {
	length := len(helper.FilesInfo)
	for i := 0; i < length/2; i++ {
		helper.FilesInfo[i], helper.FilesInfo[length-i-1] = helper.FilesInfo[length-i-1], helper.FilesInfo[i]
	}
}

// ***************************** THE -l(listing format) FLAG *************************
func Flagsl(_path string) {
	parseArg.Result = helper.GetFileInfo(_path)
}

// ***************************** WITHOUT FLAG *************************
func Flags(_path string) []helper.FilesEntry {
	parseArg.Result = []string{}
	return helper.GetFile(_path, false)
}

// ************************************* GET FILE NAME *****************************
func Getfilename(files []helper.FilesEntry) {
	for _, v := range files {
		if v.EntryFile.IsDir() {
			parseArg.Result = append(parseArg.Result, Colors["blue"]+v.P_name+Colors["reset"])
		} else {
			parseArg.Result = append(parseArg.Result, v.P_name)
		}
	}
}
