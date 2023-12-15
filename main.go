package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type DataStruct struct {
	data  []string
	count int
}

func FilePathWalkDir(root string) (DataStruct, DataStruct, error) {
	var files DataStruct
	var dirs DataStruct
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.ModTime().After(time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC)) {
			if !info.IsDir() {
				files.data = append(files.data, path)
				files.count++
				fmt.Println(path)
			} else {
				dirs.data = append(dirs.data, path)
				dirs.count++

			}
		}

		return nil
	})
	return dirs, files, err
}

func main() {

	source := flag.String("s", "", "source")
	destination := flag.String("d", "", "destination")
	flag.Parse()
	dirs, files, err := FilePathWalkDir(*source)
	if err != nil {
		panic(err)
	}

	fmt.Println("Dirs: ", dirs.count, "Files: ", files.count)

	for _, dir := range dirs.data {
		dir = strings.ReplaceAll(dir, *source, *destination)
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			panic(err)
		}
	}
	for _, file := range files.data {
		data, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		newFile := strings.ReplaceAll(file, *source, *destination)

		err = os.WriteFile(newFile, data, 0777)
		if err != nil {
			fmt.Println(err)
		}

	}

}
