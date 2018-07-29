package component

import (
	"github.com/wplib/go-lib/constant"
)
type ComponentList []*Component

type Component struct {
	class ComponentClass
	*ComponentType
	*ComponentLocation
}

func (c *Component) GetType() string {
	return c.ComponentType.GetType()
}

func (c *Component) GetLocation() string {
	return c.ComponentLocation.GetLocation()
}

func NewComponent(class ComponentClass,typestr,refstr string) *Component {
	var err error
	ct := NewComponentType()
	err = ct.Parse(typestr)
	if err != nil {
		panic(err)
	}
	cr := NewComponentLocation()
	err = cr.Parse(refstr)
	if err != nil {
		panic(err)
	}
	return &Component{
		class: class,
		ComponentType: ct,
		ComponentLocation: cr,
	}
}

//func NewStackComponent(typestr,refstr string) *Component {
//	return NewComponent(constant.StackComponent,typestr,refstr)
//}

func NewServiceComponent(typestr,refstr string) *Component {
	return NewComponent(constant.ServiceComponent,typestr,refstr)
}

//func NewExecutableComponent(typestr,refstr string) *Component {
//	return NewComponent(constant.ExecutableComponent,typestr,refstr)
//}
//
//func NewScriptComponent(typestr,refstr string) *Component {
//	return NewComponent(constant.ScriptComponent,typestr,refstr)
//}
//
//func NewSourceComponent(typestr,refstr string) *Component {
//	return NewComponent(constant.SourceComponent,typestr,refstr)
//}
//
//func NewDataComponent(typestr,refstr string) *Component {
//	return NewComponent(constant.DataComponent,typestr,refstr)
//}
//
//func NewMediaComponent(typestr,refstr string) *Component {
//	return NewComponent(constant.MediaComponent,typestr,refstr)
//}

