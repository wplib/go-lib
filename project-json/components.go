
package project_json

import "strings"

type ComponentList []*Component

type Component struct {
	Class ComponentClass
	Type  *ComponentType
	Source string
	Stack string
	Name string
	Version string
}
func NewStackComponent(n string) *Component {
	return NewComponent(StackComponent,n)
}
func NewServiceComponent(n string) *Component {
	return NewComponent(ServiceComponent,n)
}
func NewExecutableComponent(n string) *Component {
	return NewComponent(ExecutableComponent,n)
}
func NewScriptComponent(n string) *Component {
	return NewComponent(ScriptComponent,n)
}
func NewSourceComponent(n string) *Component {
	return NewComponent(SourceComponent,n)
}
func NewDataComponent(n string) *Component {
	return NewComponent(DataComponent,n)
}
func NewMediaComponent(n string) *Component {
	return NewComponent(MediaComponent,n)
}

func NewComponent(c ComponentClass, n string) *Component {
	return &Component{
		Class: c,
		Name: n,
	}
}

type ComponentClass int

const (
	StackComponent = iota
	ServiceComponent
	ExecutableComponent
	ScriptComponent
	SourceComponent
	DataComponent
	MediaComponent
)

type ComponentType struct {
	Source string
	Stack string
	Type string
	Version string
}

func NewComponentType(n string) *ComponentType {
	ct := strings.Split(n,"/")
	return &ComponentType{
		Stack: ct[0],
		Type: ct[1],
	}
}
func (ct *ComponentType) Name() string {
	return ct.Stack+"/"+ct.Type
}


