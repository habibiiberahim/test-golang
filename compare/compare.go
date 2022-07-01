package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	sourceFiles := getFiles("./source")
	targetFiles := getFiles("./target")

	for i := 0; i < len(sourceFiles); i++ {
		for _, sFile := range sourceFiles[i] {
			if !strings.Contains(strings.Join(targetFiles[i], " "), sFile) {
				fmt.Printf("%s NEW\n", sFile)
			}
		}
		for _, tFile := range targetFiles[i] {
			if !strings.Contains(strings.Join(sourceFiles[i], " "), tFile) {
				fmt.Printf("%s DELETED\n", tFile)
			}
		}
	}
}

func getDirectory(pathDir string) []string {
	var dir []string

	err := filepath.Walk(pathDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Print(err)
		}
		if info.IsDir() {
			dir = append(dir, path)
		}
		return nil
	})
	if err != nil {
		fmt.Print(err)
	}
	return dir
}

func getFiles(pathDir string) [][]string {
	dir := getDirectory(pathDir)
	files := make([][]string, len(dir))

	i := 0
	for _, d := range dir {
		file, err := ioutil.ReadDir(d)
		if err != nil {
			fmt.Print(err)
		}

		if strings.Contains(d, "./source") {
			d = strings.Replace(d, "./source", "", -1)
		} else if strings.Contains(d, "source") {
			d = strings.Replace(d, "source/", "", -1)
		} else if strings.Contains(d, "./target") {
			d = strings.Replace(d, "./target", "", -1)
		} else if strings.Contains(d, "target") {
			d = strings.Replace(d, "target/", "", -1)
		}

		for _, f := range file {
			if !f.IsDir() {
				files[i] = append(files[i], d+"/"+f.Name())
			}
		}
		i++

	}
	return files
}
