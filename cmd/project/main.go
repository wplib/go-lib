/**
 *
 */
package main

import (
	prjs "github.com/wplib/project-cli/project-json"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ERROR:", r)
		}
	}()
	pf := prjs.NewProjectFile()
	pf.SetFilepath("../project.json")
	pf.LoadJSON()
	for i, c:= range pf.GetComponents() {
		fmt.Printf("\n[%d] %-22v %v",i,c.GetType()+":",c.GetReference())
	}
}


