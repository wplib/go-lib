package project_json

type ComponentList []*Component

type Component struct {
	class ComponentClass
	*ComponentType
	*ComponentLocation
}

func (c *Component) GetType() string {
	return c.ComponentType.GetLocation()
}

func (c *Component) GetLocation() string {
	return c.ComponentLocation.GetLocation()
}

func NewComponent(class ComponentClass,typestr,refstr string) *Component {
	var err error
	ct := NewComponentType()
	//err = ct.Parse(typestr)
	//if err != nil {
	//	panic(err)
	//}
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
//	return NewComponent(StackComponent,typestr,refstr)
//}

func NewServiceComponent(typestr,refstr string) *Component {
	return NewComponent(ServiceComponent,typestr,refstr)
}

//func NewExecutableComponent(typestr,refstr string) *Component {
//	return NewComponent(ExecutableComponent,typestr,refstr)
//}
//
//func NewScriptComponent(typestr,refstr string) *Component {
//	return NewComponent(ScriptComponent,typestr,refstr)
//}
//
//func NewSourceComponent(typestr,refstr string) *Component {
//	return NewComponent(SourceComponent,typestr,refstr)
//}
//
//func NewDataComponent(typestr,refstr string) *Component {
//	return NewComponent(DataComponent,typestr,refstr)
//}
//
//func NewMediaComponent(typestr,refstr string) *Component {
//	return NewComponent(MediaComponent,typestr,refstr)
//}

