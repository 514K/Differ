package main

import (
	"fmt"
	"io"
	"os"
)

var ANSI_RESET string = "\u001B[0m"
var ANSI_RED string = "\u001b[48;5;210m"
var ANSI_GREEN string = "\u001b[42m"
var ANSI_REDL string = "\u001b[48;5;9m"

func main() {
	if len(os.Args) != 3 {
		return
	} else {
		// fmt.Printf("We have 3 args\n")
		file1, err1 := os.Open(os.Args[1])
		file2, err2 := os.Open(os.Args[2])

		if err1 != nil || err2 != nil {
			fmt.Printf("Error: %v ; %v\n", err1.Error(), err2.Error())
		} else {
			defer file1.Close()
			defer file2.Close()

			data := make([]byte, 64)

			// lines1 := make([]string, 64)
			// lines2 := make([]string, 64)

			// tmp := ""
			for {
				n, err := file1.Read(data)
				if err == io.EOF {
					break
				}
				fmt.Print(string(data[:n]))
				break
			}
		}
	}
	fmt.Printf(ANSI_RED + "My red text?" + ANSI_REDL + "Sas" + ANSI_RED + "Ke" + ANSI_RESET)
}
