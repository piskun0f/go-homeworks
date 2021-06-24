package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

const separator string = string(os.PathSeparator)
const tabulation string = "\t"
const middlePrefix string = "├───"
const middleStick string = "│"
const lastPrefix string = "└───"

func getSubs(path string, isGetFiles bool, isLasts []bool) (string, error) {
	var res string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return res, err
	}

	sort.SliceStable(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	var prefix string
	for _, isLast := range isLasts {
		if !isLast {
			prefix += middleStick + tabulation
		} else {
			prefix += tabulation
		}
	}

	if isGetFiles {
		for i, f := range files {
			isLast := i == len(files)-1

			res += prefix

			if isLast {
				res += lastPrefix
			} else {
				res += middlePrefix
			}

			res += f.Name()

			stat, err := os.Stat(path + separator + f.Name())
			fSize := stat.Size()

			if err == nil {
				if !f.IsDir() && fSize == 0 {
					res += " (empty)"
				} else if !f.IsDir() {
					res += " (" + strconv.Itoa(int(fSize)) + "b)"
				}
			}
			res += "\n"

			subs, err := getSubs(path+separator+f.Name(), isGetFiles, append(isLasts, isLast))
			if err == nil {
				res += subs
			}
		}
	} else {
		var files1 []os.FileInfo
		for _, f := range files {
			if f.IsDir() {
				files1 = append(files1, f)
			}
		}

		for i, f := range files1 {
			isLast := i == len(files1)-1

			res += prefix

			if isLast {
				res += lastPrefix
			} else {
				res += middlePrefix
			}

			res += f.Name()
			res += "\n"

			subs, err := getSubs(path+separator+f.Name(), isGetFiles, append(isLasts, isLast))
			if err == nil {
				res += subs
			}
		}
	}

	return res, nil
}

func dirTree(output io.Writer, path string, isPrintFiles bool) error {
	result, err := getSubs(path, isPrintFiles, make([]bool, 0))

	if err != nil {
		return err
	}

	result = result[:len(result)-1]

	fmt.Fprint(output, result)

	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
