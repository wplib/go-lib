
package project

import (
	"github.com/tidwall/gjson"
	"github.com/wplib/go-lib/component"
)

type ProjectWrapper interface {
	GetProject() *Project
	SetJSON(json []byte)
	GetComponents() component.ComponentList
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

func (p *Project) GetComponents() component.ComponentList {
	r:= gjson.GetBytes(p.json,"stack" )
	cl := component.ComponentList{}
	r.ForEach(func(k,v gjson.Result) bool {
		c:= component.NewServiceComponent(k.String(),v.String())
		cl = append(cl,c)
		return true
	})
	return cl
}

