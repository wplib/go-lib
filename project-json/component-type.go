package project_json

import (
	"strings"
	"strconv"
	"errors"
	"fmt"
)

/**
 * @todo Pull these from defaults somehow
 * These constants is only temporary
 */
const (
	_DefaultCTypeHost = "wplib.org"
	_DefaultCTypeStack = "wordpress"
	_DefaultCTypeVersion = "1"
)


type ComponentType struct {
	Type    string
	Host    string
	Stack   string
	Role    string
	Version byte
}

func NewComponentType() *ComponentType {
	return &ComponentType{
		Version: 1,
	}
}

func (ct *ComponentType) Parse(ts string) error {
	var e error
	ta := []string{_DefaultCTypeHost, _DefaultCTypeStack, "", _DefaultCTypeVersion}
	tp := strings.Split(ts, "/")
	tl := len(tp)
	if tl > 3 {
		msg:= fmt.Sprintf("Component type ['%v'] can only have two slashes; as in: '{host}/{stack}/{role}'.",ts)
		e = errors.New(msg)
		return e
	}
	for i := 2; i >= 0; i-- {
		tpi := 3-i-1
		tai:= 5-i-tl
		if (tpi>len(tp)-1||tai<0) {
			break
		}
		ta[tai] = tp[tpi]
	}
	tp = strings.Split(ta[2], ":")
	tl = len(tp)
	if tl > 1 {
		msg:= fmt.Sprintf("Component type ['%v'] can only have one colon (to denote version.)",ts)
		e = errors.New(msg)
		return e
	}
	if (len(tp) >= 2) {
		ta[2] = tp[0]
		ta[3] = tp[1]
	}
	vn, err := strconv.Atoi(ta[3])
	if err != nil {
		vn = 1
	}
	ct.Type = ts
	ct.Host = ta[0]
	ct.Stack = ta[1]
	ct.Role = ta[2]
	ct.Version = byte(vn)
	return nil
}

func (ct *ComponentType) FullType() string {
	return ct.Host + "/" + ct.Stack + "/" + ct.Role + ":" + strconv.Itoa(int(ct.Version))
}

