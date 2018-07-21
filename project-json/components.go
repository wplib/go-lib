
package project_json

type ComponentList []*Component

type Component struct {
	Index int
	Name  string
	Type  ComponentType
	Class ComponentClass
}

type ComponentClass int

const (
	ServiceComponent  = iota
	ExecutableComponent
	ScriptComponent
	SourceComponent
	DataComponent
	MediaComponent
)

type ComponentType struct {
	Source string
	Group string
	Type string
	Version string
}
