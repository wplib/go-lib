/**
 *
 */
package main

import (
	"fmt"
	"github.com/wplib/project-cli/project-json"
)

func main() {
	project:= project_json.NewProject()
	for _, c:= range project.GetStackComponents() {
		t:= c.Type
		fmt.Printf("\n[%d] %-22v %v",c.Index,t.Group+"/"+t.Type+":",c.Name)
	}

}
