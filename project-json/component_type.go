package project_json

import (
	"github.com/wplib/project-cli/component_type"
	"github.com/wplib/project-cli/constant"
)

type ComponentType struct {
	*Location
}

func NewComponentType() * ComponentType {
	l := NewLocation(constant.IntegerVersionStyle)
	l.SetDefaults(
		component_type.DefaultHost(),
		component_type.DefaultStack(),
		component_type.DefaultName(),
		component_type.DefaultVersion(),
	)
	return &ComponentType{
		Location: l,
	}
}

func (ct *ComponentType) GetStack() string {
	return ct.GetGroup()
}

func (ct *ComponentType) GetVersion() int {
	return ct.GetIntegerVersion()
}

func (ct *ComponentType) GetStringVersion() string {
	return ct.Location.GetVersion()
}

func (ct *ComponentType) GetType() string {
	return ct.GetHost() +
		"/" + ct.GetStack() +
		"/" + ct.GetName() +
		":" + ct.GetStringVersion()
}

