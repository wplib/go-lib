
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
	cl := ComponentList{}
	r.ForEach(func(k,v gjson.Result) bool {
		c:= NewServiceComponent(k.String(),v.String())
		cl = append(cl,c)
		return true
	})
	return cl
}

