package project_json

type ComponentList []*Component

type Component struct {
	class ComponentClass
	ctype *ComponentType
	cref  *ComponentRef
}

func (ct *Component) GetType() string {
	return ct.ctype.GetType()
}

func (ct *Component) GetReference() string {
	return ct.cref.GetLocator()
}

func NewComponent(class ComponentClass,typestr,refstr string) *Component {
	ct := NewComponentType()
	if err := ct.Parse(typestr); err != nil {
		panic(err)
	}
	cr := NewComponentRef()
	if err := cr.Parse(refstr); err != nil {
		panic(err)
	}
	return &Component{
		class: class,
		ctype: ct,
		cref:  cr,
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

