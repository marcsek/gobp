package fileapi

import (
	"fmt"
	"os"
)

type fileCreator struct {
	baseName string
}

func NewFileCreator(baseName string) *fileCreator {
	return &fileCreator{baseName: baseName}
}

func (fc *fileCreator) CreateDirs() error {
	err := os.MkdirAll(fc.baseName+"/bin", os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(fc.baseName+"/internal", os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(fc.baseName+"/cmd/main", os.ModePerm)
	if err != nil {
		return err
	}

	return err
}

func (fc *fileCreator) CreateFiles() error {
	file, err := os.Create(fc.baseName + "/cmd/main/main.go")

	if err != nil {
		return err
	}

	file.WriteString(mainBp)
	file.Close()

	file, err = os.Create(fc.baseName + "/Makefile")

	if err != nil {
		return err
	}

	file.WriteString(fmt.Sprintf(makeBp, fc.baseName))

	return err
}

const mainBp = `package main

import "fmt"

func main() {
  fmt.Print("jak se vede")
}
`

const makeBp = `build:
	go build -o bin/%s cmd/main/main.go

run:
	go run cmd/main/main.go`
