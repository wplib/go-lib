
package project_json

import (
	"github.com/tidwall/gjson"
)

type ProjectWrapper interface {
	GetProject() *Project
	SetJSON(json []byte)
	GetComponents() ComponentList
}

type Project struct {
	json []byte
	Stack ProjectStack
}

func NewProject() *Project {
	return &Project{}
}

func (p *Project) GetProject() *Project {
	return p
}

func (p *Project) SetJSON(json []byte) {
	p.json=json
}

func (p *Project) GetComponents() ComponentList {
	r:= gjson.GetBytes(p.json,"stack" )
	cl := make(ComponentList,len(r.Map()))
	index := 0
	r.ForEach(func(k,v gjson.Result) bool {
		cl[index] = NewServiceComponent(k.String(),v.String())
		index++
		return true
	})
	return cl
}

