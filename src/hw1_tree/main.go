package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func print(s string) {
	fmt.Println(s)
}

func getFiles(path string, depth int) (string, error) {
	var res string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return res, err
	}

	var buf []string = make([]string, 0, len(files))

	for _, f := range files {
		buf = append(buf, f.Name())
	}

	sort.Strings(buf)

	for i, el := range buf {
		subfiles, err := getFiles(path+string(os.PathSeparator)+el, depth+1)

		if depth != 0 {
			res += strings.Repeat("|\t", depth)
		}

		if err == nil {
			res += "├───" + el + "\n" + subfiles
		} else if i == len(buf)-1 {
			res += "└───" + el + "\n"
		} else {
			res += "├───" + el + "\n"
		}
	}

	return res, nil
}

func dirTree(output io.Writer, path string, isPrintFiles bool) error {
	var result string

	fmt.Fprint(output, result)

	return nil
}

func main() {
	s, _ := getFiles("testdata", 0)
	fmt.Println(s)
	// out := os.Stdout
	// if !(len(os.Args) == 2 || len(os.Args) == 3) {
	// 	panic("usage go run main.go . [-f]")
	// }
	// path := os.Args[1]
	// printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	// err := dirTree(out, path, printFiles)
	// if err != nil {
	// 	panic(err.Error())
	// }
}
