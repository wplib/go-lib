
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
	return Project{}
}

func (p *Project) LoadJSON() {
	json, err := ioutil.ReadFile("../project.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	p.json = json
}

func (p *Project) GetStackComponents() ComponentList {
	r:= gjson.GetBytes(p.json,"stack" )
	rm:= r.Map()
	cl := make(ComponentList,len(rm))
	for n,v:= range rm {
		ct := strings.Split(n,"/")
		cl[v.Index] = &Component{
			Index: v.Index,
			Name: v.String(),
			Class: ServiceComponent,
			Type: ComponentType{
				Group: ct[0],
				Type: ct[1],
			},
		}
	}
	return cl
}
