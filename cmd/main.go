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
	for _, c:= range p.GetStackComponents() {
		t:= c.Type
		fmt.Printf("\n[%d] %-22v %v",c.Index,t.Group+"/"+t.Type+":",c.Name)
	}

}
