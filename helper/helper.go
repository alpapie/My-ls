package helper

import (
	"fmt"
	"io/fs"
	"my_ls_1/parseArg"
	"my_ls_1/verif"
	"os"
	"os/user"
	"strings"
	"syscall"
)

type FilesEntry struct{
	EntryFile fs.DirEntry
	P_name string
}
var FilesInfo = []FilesEntry{}
var Totalsize =0
var Colors = map[string]string{
	"blue":  "\033[38;2;0;122;204m",
	"linkcolor":"\033[38;2;41;155;198m",
	"reset": "\033[0m",
}
var LenCols=[7]int{}

// ***************************** RECUP THE FILE ****************************
func GetFile(path string, hidden bool) []FilesEntry {
	tab := []FilesEntry{}
	files, err := os.ReadDir(path)
	
	verif.Errorstr(err, "")
	for _, file := range files {
		_file := FilesEntry{}
		if file.IsDir() {
			if !hidden {
				if string(file.Name()[0]) != "." {
					_file.P_name=file.Name()
					_file.EntryFile=file
					tab = append(tab, _file)
				}
			} else {
				_file.P_name=file.Name()
				_file.EntryFile=file
				tab = append(tab, _file)
			}
		} else {
			if !hidden {
				if string(file.Name()[0]) != "." {
					_file.P_name=file.Name()
				_file.EntryFile=file
				tab = append(tab, _file)
				}
			} else {
				_file.P_name=file.Name()
				_file.EntryFile=file
				tab = append(tab, _file)
			}
		}
	}
	return tab
}

// *************************RECUP file INFORMATION ***************************
func GetFileInfo(_path string) []string {
	tab := [][]string{}
	//strtab := [][]string{}
	for _, v := range FilesInfo {
		for _, y := range parseArg.Result {
			_tab := []string{}
			direc := ""
			if v.EntryFile.IsDir() {
				direc = Colors["blue"] + v.P_name + Colors["reset"]
			} else {
				direc = v.P_name
			}
			_tab = _GetFileInfo(_path,y, direc, v)
			if len(_tab)>0{
				tab = append(tab, _tab)
			}
		}
	}
	
	return AlignColumb(tab)
}

func _GetFileInfo(_path,y , direc string, v FilesEntry) []string {
	tab:=[]string{}
	if y == direc {
		entries, err := v.EntryFile.Info()
		verif.Errorstr(err, "")
		date := entries.ModTime().Format("Jan _2 15:04")
		mod := entries.Mode().String()
		infouser := entries.Sys()
		username := ""
		link := 0
		group := ""
		size := 0
		if stat, ok := infouser.(*syscall.Stat_t); ok {
			uid := stat.Uid
			size = int(stat.Size)
			Totalsize+=int((stat.Size+4096-1)/4096*(4096/1024))
			link = int(stat.Nlink)
			u, err := user.LookupId(fmt.Sprintf("%d", uid))
			verif.Errorstr(err, "")
			username = u.Username

			uu, err := user.LookupGroupId(fmt.Sprintf("%d", uid))
			verif.Errorstr(err, "")
			group = uu.Name
		}
		if mod[0]=='L'{
			linkk,errrr:=os.Readlink(_path+"/"+y)
			verif.Errorstr(errrr, "")
			linkk=Colors["blue"]+linkk+Colors["reset"]
			y= Colors["linkcolor"]+y+Colors["reset"]+" -> \033[1m"+linkk+"\033[0m"
		}
		tab= []string{mod,fmt.Sprint(link),fmt.Sprint(username),fmt.Sprint(group),fmt.Sprint(size),fmt.Sprint(date),"\033[1m"+y+"\033[0m"}
	}
	return tab
}

// ************************************** FORMAT THE STRING FOR ALIGN OF COLUMB *********************************
func AlignColumb(tab [][]string) []string{
	GetMaxe(tab)
	tabStrs:=[]string{}
	for _, v := range tab {
		s:=""
		for i, y := range v {
			if i > 0 && i < len(v)-1{
				s+=" "+ strings.Repeat(" ",LenCols[i]-len(y))+y
			}else if i==len(v)-1{
				s+=" "+y
			}else{
				s+=y
			}
		}
		tabStrs=append(tabStrs, s)
	}
	return tabStrs
}

//**************************************GET THE MAX FOR ALL COLUM******************************
func GetMaxe(tab [][]string){
	for i := 0; i < len(LenCols); i++ {
		for j := 0; j < len(tab); j++ {
			l:=len(tab[j][i])
			if LenCols[i]<l{
				LenCols[i]=l
			}
		}
	}
}