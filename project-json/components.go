
package project_json

import (
	"strings"
	"strconv"
)

type ComponentList []*Component

type Component struct {
	Class     ComponentClass
	Type      *ComponentType
	Reference *ComponentReference
}

func NewStackComponent(t string,r string) *Component {
	return NewComponent(StackComponent,t,r)
}

func NewServiceComponent(t string,r string) *Component {
	return NewComponent(ServiceComponent,t,r)
}

func NewExecutableComponent(t string,r string) *Component {
	return NewComponent(ExecutableComponent,t,r)
}

func NewScriptComponent(t string,r string) *Component {
	return NewComponent(ScriptComponent,t,r)
}

func NewSourceComponent(t string,r string) *Component {
	return NewComponent(SourceComponent,t,r)
}

func NewDataComponent(t string,r string) *Component {
	return NewComponent(DataComponent,t,r)
}

func NewMediaComponent(t string,r string) *Component {
	return NewComponent(MediaComponent,t,r)
}

func NewComponent(c ComponentClass,t string,r string) *Component {
	return &Component{
		Class: c,
		Type: NewComponentType(t),
		Reference: NewComponentReference(r),
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
	Type 	string
	Host    string
	Stack   string
	Role    string
	Version byte
}

/**
 * @todo Initialize missing from .defaults.componentType
 */
func NewComponentType(ts string) *ComponentType {
	ta := []string{"wplib.org","wordpress","","1"}
	tp := strings.Split(ts,"/")
	for i:=2; i>=0; i++ {
		if (len(tp)<i) {
			continue
		}
		ta[i] = tp[2-i]
	}
	tp = strings.Split(ta[2],":")
	if (len(tp)>=2) {
		ta[2] = tp[0]
		ta[3] = tp[1]
	}
	vn, err:= strconv.Atoi(ta[3])
	if err != nil {
		vn = 1
	}
	return &ComponentType{
		Type:    ts,
		Host:    ta[0],
		Stack:   ta[1],
		Role:    ta[2],
		Version: byte(vn),
	}
}
func (ct *ComponentType) FullType() string {
	return ct.Host+"/"+ct.Stack+"/"+ct.Role+":"+strconv.Itoa(int(ct.Version))
}

type ComponentReference struct {
	Host    string
	Group   string
	Name    string
	Version *ComponentVersion
}

/**
 * @todo Initialize missing from .defaults.componentType
 */
func NewComponentReference(rs string) *ComponentReference {
	ra := []string{"","","",""}
	rp := strings.Split(rs,"/")
	for i:=2; i>=0; i++ {
		if (len(rp)<i) {
			continue
		}
		ra[i] = rp[2-i]
	}
	rp = strings.Split(ra[2],":")
	if (len(rp)>=2) {
		ra[2] = rp[0]
		ra[3] = rp[1]
	}

	return &ComponentReference{
		Host:    ra[0],
		Group:   ra[1],
		Name:    ra[2],
		Version: NewComponentVersion(ra[3]),
	}
}

type ComponentVersion struct {
	Version string
	Major   byte
	Minor   byte
	Patch   byte
}

/**
 * @todo Initialize missing from .defaults.componentType
 */
func NewComponentVersion(sv string) *ComponentVersion {
	va := []byte{0, 0, 0}
	vp := strings.Split(sv,".")
	for i:=0; i<=2; i++ {
		if (len(vp)<i) {
			continue
		}
		vn, err:= strconv.Atoi(vp[i])
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
