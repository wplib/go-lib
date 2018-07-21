
package project_json

type ComponentList []*Component

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
