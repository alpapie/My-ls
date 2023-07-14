package verif

import (
	"fmt"
	"os"
	"strings"
)

var GoodFlag = "lRart"

// ************************** CHECK IF FLAG IS WELL FORMED ************************
func IsAFlag(s string) bool {
	if len(s) > 1 && s[0] == '-' {
		for _, v := range s[1:] {
			if !strings.Contains(GoodFlag, string(v)) {
				return false
			}
		}
		return true
	}
	return false
}

// ******************ERROR MANAGE ****************************************************
func Errorstr(err error, er string) {
	if err != nil {
		fmt.Println("\033[31m", err, "\033[0m")
		os.Exit(0)
	} else if er != "" {
		fmt.Println("\033[31m", er, "\033[0m")
		os.Exit(0)
	}
}
