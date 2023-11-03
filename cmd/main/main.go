package main

import (
	"fmt"
	"github.com/marcsek/gobp/internal/file_api"
	"log"
	"os/exec"
)

func main() {
	var baseName string

	for len(baseName) < 1 {
		fmt.Print("Name of your project: ")
		fmt.Scanln(&baseName)

		if len(baseName) < 1 {
			fmt.Println("You need to specify project name")
		}
	}

	projectName := fmt.Sprintf("github.com/marcsek/%s", baseName)
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = "./" + baseName

	fc := fileapi.NewFileCreator(baseName)
	err := fc.CreateDirs()

	if err != nil {
		log.Fatal(err)
	}

	err = fc.CreateFiles()

	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
