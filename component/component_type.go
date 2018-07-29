package component

import (
	"github.com/wplib/go-lib/defaults"
	"github.com/wplib/go-lib/constant"
	"github.com/wplib/go-lib/location"
)

type ComponentType struct {
	*location.Location
}

func NewComponentType() * ComponentType {
	l := location.NewLocation(constant.IntegerVersionStyle)
	l.SetDefaults(
		defaults.DefaultTypeHost(),
		defaults.DefaultTypeStack(),
		defaults.DefaultTypeName(),
		defaults.DefaultTypeVersion(),
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

