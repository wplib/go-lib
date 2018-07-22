/**
 *
 */
package main

import (
	"fmt"
	"github.com/wplib/project-cli/project-json"
)

func main() {
	p := project_json.NewProject()
	p.LoadJSON()
	for i, c:= range p.GetComponents() {
		t:= c.Type
		fmt.Printf("\n[%d] %-22v %v",i,t.Name()+":",c.Name)
	}

}
