package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var ANSI_RESET string = "\u001B[0m"
var ANSI_RED string = "\u001b[48;5;210m"
var ANSI_GREEN string = "\u001b[42m"
var ANSI_REDL string = "\u001b[48;5;9m"

func main() {
	t := time.Now()
	if len(os.Args) != 3 {
		return
	} else {
		c1 := make(chan int)
		arr1 := make([]string, 0)

		c2 := make(chan int)
		arr2 := make([]string, 0)

		go readFile(os.Args[1], &arr1, c1)
		go readFile(os.Args[2], &arr2, c2)

		for {
			if <-c1 == 0 && <-c2 == 0 {
				fmt.Printf("%v\n", arr1)
				fmt.Printf("%v\n", arr2)
				break
			}
		}

		fmt.Printf("%v; %v\n", len(arr1), len(arr2))

		for i := 0; i < 3; i++ {
			for j := 0; j < 10; j++ {
				i, j = compareLines(i, j, arr1, arr2)
			}
		}
	}

	// fmt.Printf(ANSI_RED + "My red text?" + ANSI_REDL + "Sas" + ANSI_RED + "Ke" + ANSI_RESET) ШАБЛОН КОДА ДЛЯ ВЫДЕЛЕНИЯ СТРОКИ В КОНСОЛИ
	fmt.Print("\n")
	fmt.Print(time.Since(t))
}

// Будет возвращать i, j строк с которых надо продолжать поиск
func compareLines(nl1 int, nl2 int, arr1 []string, arr2 []string) (int, int) {
	// Поиск слов по строкам с номера строки nl1 из arr1 и номера строки nl2 из arr2
	for i := nl1; i < len(arr1); i++ {
		for j := nl2; j < len(arr2); j++ {
			// Тут сам поиск и изменение nl1 и nl2
		}
	}
	return nl1, nl2
}

func readFile(filename string, arr *[]string, ch chan int) {

	file, err := os.Open(filename)

	if err == nil {
		defer file.Close()

		data := make([]byte, 1)

		tmpstr := ""

		for {
			_, err := file.Read(data)
			if err == io.EOF {
				*arr = append(*arr, tmpstr)
				break
			}

			if string(data) != "\n" {
				tmpstr += string(data)
			} else {
				*arr = append(*arr, tmpstr+"\n")
				tmpstr = ""
			}
		}
	}
	close(ch)
}
