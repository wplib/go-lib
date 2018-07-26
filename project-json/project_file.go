package project_json

import (
	"io/ioutil"
	"fmt"
	"os"
)

type ProjectFile struct {
	filepath string
	ProjectWrapper
}

func NewProjectFile() *ProjectFile {
	return &ProjectFile{
		filepath: "./project.json",
		ProjectWrapper: NewProject(),
	}
}

func (pf *ProjectFile) LoadJSON() {
	json, err := ioutil.ReadFile(pf.filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	pf.SetJSON(json)
}

func (pf *ProjectFile) SetFilepath(filepath string) {
	pf.filepath = filepath
}

