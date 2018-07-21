
package project_json

import (
	"io/ioutil"
	"fmt"
	"os"
	"github.com/mikeschinkel/gjson"
	"strings"
)

type Project struct {
	json []byte
	Stack Stack
}

func NewProject() Project {
	var project Project

	json, err := ioutil.ReadFile("../project.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	project.json = json
	return project
}

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
