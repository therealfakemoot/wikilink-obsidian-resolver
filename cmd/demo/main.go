package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var vaultRoot = "testdata/test_vault/"

func main() {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("couldn't get current working directory before opening vault dir: %s\n", err)
		return
	}

	absPath := filepath.Join(wd, vaultRoot)
	vaultFS := os.DirFS(absPath)

	files, err := fs.Glob(vaultFS, "*")
	if err != nil {
		fmt.Printf("couldn't run glob: %s\n", err)
		return
	}
	fmt.Println(files)
}
