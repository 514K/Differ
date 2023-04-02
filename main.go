package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var ANSI_RESET string = "\u001B[0m"
var ANSI_RED string = "\u001b[48;5;210m"
var ANSI_GREEN string = "\u001b[42m"
var ANSI_REDL string = "\u001b[48;5;9m"

func main() {
	// t := time.Now()
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
				break
			}
		}

		compareLines(0, 0, arr1, arr2)
	}

	// fmt.Printf(ANSI_RED + "My red text?" + ANSI_REDL + "Sas" + ANSI_RED + "Ke" + ANSI_RESET) ШАБЛОН КОДА ДЛЯ ВЫДЕЛЕНИЯ СТРОКИ В КОНСОЛИ
	fmt.Print("\n")
	// fmt.Print(time.Since(t))
}

// Будет возвращать i, j строк с которых надо продолжать поиск
func compareLines(nl1 int, nl2 int, arr1 []string, arr2 []string) {
	// Поиск слов по строкам с номера строки nl1 из arr1 и номера строки nl2 из arr2
	strFind := false
	for i := nl1; i < len(arr1); i++ {
		// if i > 30 {

		// 	break
		// }

		if arr1[i][0] == 13 {
			continue
		}
		strFind = false
		for j := nl2; j < len(arr2); j++ {

			if arr2[j][0] == 13 {
				continue
			}

			// if j > 30 {
			// 	fmt.Printf("%v %v\n", arr1[i], i)
			// 	fmt.Printf("%v %v\n", arr2[j], j)
			// 	break
			// }

			arr1[i] = strings.TrimLeft(arr1[i], " ")
			arr2[j] = strings.TrimLeft(arr2[j], " ")

			// Тут сам поиск и изменение nl1 и nl2
			if arr1[i] != arr2[j] {
				tmpstr1 := ""
				tmpstr2 := ""
				for _, word1 := range arr1[i] {

					if string(word1) != " " && string(word1) != "\n" && string(word1) != ">" && string(word1) != string(arr1[i][len(arr1[i])-1]) {
						tmpstr1 += string(word1)
						continue
					}
					tmpstr1 += string(word1)

					for _, word2 := range arr2[j] {

						// if string(word2) != " " {
						// 	tmpstr2 += string(word2)
						// 	continue
						// }
						// tmpstr2 += string(word2)

						if string(word2) != " " && string(word2) != "\n" && string(word2) != ">" && string(word2) != string(arr2[j][len(arr2[j])-1]) {
							tmpstr2 += string(word2)
							continue
						}
						tmpstr2 += string(word2)

						// fmt.Printf("%v%v\n", tmpstr1, tmpstr2)
						// fmt.Printf("%v\n", tmpstr1 == tmpstr2)

						if tmpstr1 == tmpstr2 {
							// Строка найдена
							strFind = true
							// Запоминаем номера строк???

							if nl2 < j {
								printGreen(nl2, j, arr2)
							}
							printDifference(arr1[i], arr2[j])

							// Дальше надо выйти из цикла по j
							nl2 = j + 1
						}

						tmpstr1 = ""
						tmpstr2 = ""

						if strFind {
							break
						}
					}
					if strFind {
						break
					}
				}

				if strFind {
					break
				}

			} else {

				// if nl2 < j {
				// 	printGreen(nl2-1, j, arr2)
				// }
				fmt.Print(arr2[j])

				nl2 = j + 1
				strFind = true
			}

			if strFind {
				break
			}

		}
	}
	// return nl1, nl2
}

func printGreen(lineStart int, lineStop int, arr2 []string) {
	for i := lineStart; i < lineStop; i++ {
		fmt.Print(ANSI_GREEN + arr2[i] + ANSI_RESET)
	}
}

func printDifference(str1 string, str2 string) {
	tmpstr1 := ""
	tmpstr2 := ""

	curWord := 0

	for _, word1 := range str1 {
		if string(word1) != " " && string(word1) != "\n" && string(word1) != ">" && string(word1) != string(str1[len(str1)-1]) {
			tmpstr1 += string(word1)
			continue
		}
		tmpstr1 += string(word1)

		for j := curWord; j < len(str2); j++ {
			if string(str2[j]) != " " && string(str2[j]) != "\n" && string(str2[j]) != ">" && string(str2[j]) != string(str2[len(str2)-1]) {
				tmpstr2 += string(str2[j])
				continue
			}
			tmpstr2 += string(str2[j])

			if tmpstr1 != tmpstr2 {
				fmt.Print(ANSI_GREEN + tmpstr2 + ANSI_RESET)
			} else {
				fmt.Print(tmpstr2)
			}

			tmpstr1 = ""
			tmpstr2 = ""
			curWord = j + 1
			break
		}
	}
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
