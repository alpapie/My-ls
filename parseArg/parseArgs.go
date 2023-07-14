package parseArg

import "my_ls_1/verif"

var Flags = map[string]bool{
	"":true,
	"a":false,
	"r":false,
	"l":false,
	"R":false,
	"t":false,
}
var PathFolder = []string{}
var Result=[]string{}

// *********************** PARSE THE ARGUMENT DATA BY GETTING THE FLAG AND THE PATH FOLDER*********
func ParseArgs(a []string) {
	for _, v := range a {
		if v[0] == '-' {
			if !verif.IsAFlag(v){
				verif.Errorstr(nil,"error in the flag")
			}
			ParseFlags(v[1:])
		} else {
			PathFolder = append(PathFolder, v)
		}
	}
}

// *********************** PARSE THE STRING FLAGS *********
func ParseFlags(a string) {
	for _, v := range a {
		Flags[string(v)]=true
	}
}
