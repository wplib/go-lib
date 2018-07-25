package project_json

import (
	"strings"
	"fmt"
	"errors"
	"strconv"
)

const (
	IntegerVersionStyle = iota
	DottedVersionStyle
)

type LocationDefaults [4]string

type Location struct {
	raw      string
	defaults LocationDefaults
	host     string
	group    string
	name     string
	version  string
	style    byte
	dotted   *DottedVersionable
	integer  int
}

func NewLocation(vs byte) *Location {
	return &Location{
		style: vs,
	}
}

func (l *Location) GetLocation() string {
	return l.host + "/" + l.group + "/" + l.name + ":" + l.version
}

func (l *Location) SetDefaults(host, group, item, version string) *Location {
	if (l.style == IntegerVersionStyle) {
		vn, err := strconv.Atoi(version)
		if err != nil {
			msg := fmt.Sprintf( "Invalid default value ['%v'] for integer-style version: err", version, err )
			err = errors.New(msg)
			panic(err)
		}
		version = strconv.Itoa(vn)
	} else if l.style == DottedVersionStyle {
		dv, err := l.parseDottedVersion(version)
		if err != nil {
			panic(err)
		}
		version = dv.GetVersion()
	}
	l.defaults = LocationDefaults{host, group, item, version}
	return l
}

func (l *Location) GetRawLocation() string {
	return l.raw
}

func (l *Location) GetDefaults() LocationDefaults {
	return l.defaults
}

func (l *Location) GetHost() string {
	return l.host
}

func (l *Location) GetGroup() string {
	return l.group
}

func (l *Location) GetName() string {
	return l.name
}

func (l *Location) GetVersion() string {
	return l.version
}

func (l *Location) GetDottedVersionable() *DottedVersionable {
	return l.dotted
}

func (l *Location) GetDottedVersion() string {
	return l.dotted.GetVersion()
}

func (l *Location) GetIntegerVersion() int {
	return l.integer
}

func (l *Location) Parse(locstr string) error {
	la := l.defaults
	lp := strings.Split(locstr, "/")
	ll := len(lp)
	if ll > 3 {
		msg := fmt.Sprintf("Location ['%v'] can only have two slashes; as in: '{host}/{stack}/{role}'.", locstr)
		err := errors.New(msg)
		return err
	}
	for i := 2; i >= 0; i-- {
		tpi := 3 - i - 1
		tai := 5 - i - ll
		if (tpi > len(lp)-1 || tai < 0) {
			break
		}
		la[tai] = lp[tpi]
	}
	lp = strings.Split(la[2], ":")
	ll = len(lp)
	if ll > 2 {
		msg := fmt.Sprintf("Location ['%v'] can only have one colon (to denote version.)", locstr)
		err := errors.New(msg)
		return err
	}
	if ll == 2 {
		la[2] = lp[0]
		la[3] = lp[1]
	}
	l.raw = locstr
	l.host = la[0]
	l.group = la[1]
	l.name = la[2]
	var err error
	if l.style == IntegerVersionStyle {
		l.integer, err = l.parseIntegerVersion(la[3])
		if err != nil {
			return err
		}
		l.version = strconv.Itoa(l.integer)

	} else if l.style == DottedVersionStyle {
		l.dotted, err = l.parseDottedVersion(la[3])
		if err != nil {
			return err
		}
		l.version = l.dotted.GetVersion()
	}
	return nil
}

func (l *Location) parseIntegerVersion(verstr string) (int,error) {
	iv, err := strconv.Atoi(verstr)
	if err != nil {
		msg := "Version ['%v'] in locator ['%v'] is not a valid integer: %v"
		msg = fmt.Sprintf( msg, verstr, l.raw, err )
		err = errors.New(msg)
		return 0,err
	}
	return iv,nil
}

func (l *Location) parseDottedVersion(verstr string) (*DottedVersionable,error) {
	dv := NewDottedVersionable()
	err := dv.Parse(verstr)
	if err != nil {
		msg := "Invalid value ['%v'] in locator ['%v'] for dotted-style version: %v"
		msg = fmt.Sprintf( msg, verstr, l.raw, err )
		err = errors.New(msg)
		return nil,err
	}
	return dv,nil
}

