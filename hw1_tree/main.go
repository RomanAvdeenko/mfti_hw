package main

import (
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"sort"
)

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

func dirTree(out io.Writer, path string, printFiles bool) error {
	var nestingLevels = make([]bool, 0, 4)
	return dirTreeCore(out, path, printFiles, 0, &nestingLevels)
}

func dirTreeCore(out io.Writer, path string, printFiles bool, currentNesting int, nestingLevels *[]bool) error {
	// Get directory files
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	// Remove files if no "-f" environment parameter
	if !printFiles {
		dirs := make([]fs.FileInfo, 0)
		for _, file := range files {
			if file.IsDir() {
				dirs = append(dirs, file)
			}
		}
		files = dirs
	}
	// Sorting
	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	// Writing out
	for i := 0; i < len(files); i++ {
		file := files[i]
		isLast := i == len(files)-1
		// Increase slice size if needed
		if currentNesting+1 > len(*nestingLevels) {
			*nestingLevels = append(*nestingLevels, isLast)
		} else {
			(*nestingLevels)[currentNesting] = isLast
		}
		// ...tab symbols
		print("\t")
		for n := 0; n < currentNesting; n++ {
			if (*nestingLevels)[n] {
				print(" ")
			} else {
				print("│")
			}
			print("\t")
		}
		// ...file names
		if isLast {
			print("└───")
		} else {
			print("├───")
		}
		print(file.Name())
		if !file.IsDir() {
			print(" (")
			if fileSize := file.Size(); fileSize == 0 {
				print("empty")
			} else {
				print(fileSize, "b")
			}
			print(")")
		}
		println()

		// Call recursively if it's a folder
		if file.IsDir() {
			err = dirTreeCore(out, path+string(os.PathSeparator)+file.Name(), printFiles, currentNesting+1, nestingLevels)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
