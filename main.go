/**
 * See also: https://github.com/tidwall/gjson
 */
package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"github.com/mikeschinkel/gjson"
)

type Project struct {
	json []byte
	Stack Stack
}

func NewProject() Project {
	var project Project

	json, err := ioutil.ReadFile("./project.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	project.json = json
	return project
}

type ComponentList []*Component

func (p *Project) GetStackComponents() ComponentList {
	r:= gjson.GetBytes(p.json,"stack" )
	rm:= r.Map()
	cl := make(ComponentList,len(rm))
	for n,v:= range rm {
		ct := strings.Split(n,"/")
		cl[v.Index] = &Component{
			Index: v.Index,
			Name:  v.String(),
			Type: ComponentType{
				Group: ct[0],
				Type: ct[1],
			},
		}
	}
	return cl
}

type Stack struct {
	Components ComponentList
}

type Component struct {
	Index int
	Name  string
	Type  ComponentType
	Class ComponentClass
}

type ComponentClass struct {
	Source string
	Group string
	Type string
	Version string
}

type ComponentType struct {
	Source string
	Group string
	Type string
	Version string
}

func main() {
	project:= NewProject()
	for _, c:= range project.GetStackComponents() {
		t:= c.Type
		fmt.Printf("\n[%d] %-22v %v",c.Index,t.Group+"/"+t.Type+":",c.Name)
	}

}
