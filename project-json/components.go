
package project_json

import (
	"strings"
	"strconv"
)

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
	Host    string
	Stack   string
	Role    string
	Version byte
}

/**
 * @todo Initialize missing from .defaults.componentType
 */
func NewComponentType(n string) *ComponentType {
	ct := strings.Split(n,"/")
	var c *ComponentType
	switch(len(ct)) {
	case 1:
		c= &ComponentType{
			Host: "wplib.org",
			Stack: "wordpress",
			Role:  ct[0],
		}
	case 2:
		c= &ComponentType{
			Host: "wplib.org",
			Stack: ct[0],
			Role:  ct[1],
		}
	case 3:
		c= &ComponentType{
			Host: ct[0],
			Stack: ct[1],
			Role:  ct[2],
		}
	}
	return c
}
func (ct *ComponentType) Name() string {
	return ct.Stack+"/"+ct.Role
}


type ComponentVersion struct {
	Version string
	Major   byte
	Minor   byte
	Patch   byte
}
func NewComponentVersion(sv string) *ComponentVersion {
	va := []byte{0, 0, 0}
	vp := strings.Split(sv,".")
	for i:=0; i<=2; i++ {
		if (len(vp)<i) {
			continue
		}
		vn, err:= strconv.Atoi(vp[0])
		if err != nil {
			break
		}
		va[i] = byte(vn)
	}
	return &ComponentVersion{
		Version: sv,
		Major:   va[0],
		Minor:   va[1],
		Patch:   va[2],
	}
}
func (cv *ComponentVersion) FullVersion() string {
	return strconv.Itoa(int(cv.Major))+"."+
		strconv.Itoa(int(cv.Minor))+"."+
		strconv.Itoa(int(cv.Patch))
}
