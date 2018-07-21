
package project_json

import (
	"io/ioutil"
	"fmt"
	"os"
	"github.com/mikeschinkel/gjson"

)

type Project struct {
	json []byte
	Stack ProjectStack
}

func NewProject() *Project {
	return &Project{}
}

func (p *Project) LoadJSON() {
	json, err := ioutil.ReadFile("../project.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	p.json = json
}

func (p *Project) GetComponents() ComponentList {
	r:= gjson.GetBytes(p.json,"stack" )
	rm:= r.Map()
	cl := make(ComponentList,len(rm))
	for n,v:= range rm {
		sc := NewServiceComponent(v.String())
		sc.Type = NewComponentType(n)
		sc.Index = v.Index
		cl[sc.Index] = sc
	}
	return cl
}

type ProjectStack struct {
	Components ComponentList
}
