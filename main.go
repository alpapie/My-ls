package main

import (
	"fmt"
	"io/fs"
	"my_ls_1/ls"
	"my_ls_1/parseArg"
	"my_ls_1/verif"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		parseArg.ParseArgs(args)
		if len(parseArg.PathFolder) == 0 {
			parseArg.PathFolder = append(parseArg.PathFolder, ".")
		}
		for _, v := range parseArg.PathFolder {
			// ***************************** THE -R(recursive) FLAG *************************
			if parseArg.Flags["R"] {
				fileSystem := os.DirFS(v)
				fs.WalkDir(fileSystem, ".", func(pathfile string, p fs.DirEntry, err error) error {
					verif.Errorstr(err, "")
					if p.IsDir() {
						if len(strings.Split(v, "/"))>1 {
							fmt.Println(v+pathfile, ":")
							ls.Resolve(v + pathfile)
						} else {
							fmt.Println(v+"/"+pathfile, ":")
							ls.Resolve(v +"/"+ pathfile)
						}
					}
					return err
				})
			} else {
				if len(parseArg.PathFolder)>1{
					fmt.Println(v, ":")
				}
				ls.Resolve(v)
				// fmt.Println()
			}
		}
	} else {
		ls.Resolve(".")
	}
}

// fs.WalkDir(f, ".", func(path string, entry fs.DirEntry, err error) error {
// 	if entry.IsDir() {
// 		if !strings.HasPrefix(entry.Name(), ".") || (strings.HasPrefix(entry.Name(), ".") && (flags.ShowAll || len(entry.Name()) == 1)) {
// 			if entry.Name() != "." {
// 				fmt.Println(dirPath + "/" + path + ":")
// 			} else {
// 				fmt.Println(dirPath + ":")
// 			}
// 			my_ls.ListFiles(dirPath+"/"+path, flags)
// 			fmt.Println()
// 		} else {
// 			err = fs.SkipDir
// 		}
// 	}
// 	return err
// })
