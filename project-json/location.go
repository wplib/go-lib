package project_json

import (
	"strings"
	"fmt"
	"errors"
	"strconv"
	"runtime"
	"github.com/wplib/project-cli/constant"
)

type LocationDefaults [4]string

type Location struct {
	parsed   bool
	raw      string
	defaults LocationDefaults
	host     string
	group    string
	name     string
	version  string
	style    byte
	*DottedVersion
	integer  byte
}

func NewLocation(vs byte) *Location {
	return &Location{
		style:  vs,
		parsed: false,
	}
}

func (l *Location) SetDefaults(host, group, item, version string) *Location {
	if (l.style == constant.IntegerVersionStyle) {
		vn, err := strconv.Atoi(version)
		if err != nil {
			msg := fmt.Sprintf("Invalid default value ['%v'] for integer-style version: %v", version, err)
			err = errors.New(msg)
			panic(err)
		}
		version = strconv.Itoa(vn)
	} else if l.style == constant.DottedVersionStyle {
		dv, err := l.parseDottedVersion(version)
		if err != nil {
			panic(err)
		}
		version = dv.GetVersion()
	}
	l.parsed = false
	l.defaults = LocationDefaults{host, group, item, version}
	return l
}

func chkParsed(l *Location) {
	if ! l.parsed {
		// See: https://stackoverflow.com/a/25927915/102699
		// See also: https://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go
		// See also: https://lawlessguy.wordpress.com/2016/04/17/display-file-function-and-line-number-in-go-golang/
		pc := make([]uintptr, 1)
		runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[0])
		file, line := f.FileLine(pc[0])
		msg := "Parse() not yet called, in %s:%d"
		msg = fmt.Sprintf(msg, file, line)
		panic(errors.New(msg))
	}
}

func (l *Location) GetLocation() string {
	chkParsed(l)
	return l.host + "/" + l.group + "/" + l.name + ":" + l.version
}

func (l *Location) GetRawLocation() string {
	chkParsed(l)
	return l.raw
}

func (l *Location) GetParsed() bool {
	return l.parsed
}

func (l *Location) GetDefaults() LocationDefaults {
	chkParsed(l)
	return l.defaults
}

func (l *Location) GetHost() string {
	chkParsed(l)
	return l.host
}

func (l *Location) GetGroup() string {
	chkParsed(l)
	return l.group
}

func (l *Location) GetName() string {
	chkParsed(l)
	return l.name
}

func (l *Location) GetVersion() string {
	chkParsed(l)
	return l.version
}

func (l *Location) GetVersionStyle() byte {
	return l.style
}

func (l *Location) GetDottedVersion() *DottedVersion {
	chkParsed(l)
	return l.DottedVersion
}

func (l *Location) GetIntegerVersion() int {
	chkParsed(l)
	return int(l.integer)
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
	if ll < 2 {
		if l.style==constant.IntegerVersionStyle {
			la[3] = constant.DefaultIntegerVersion
		} else {
			la[3] = constant.DefaultDottedVersion
		}
	}
	if ll == 2 {
		la[2] = lp[0]
		la[3] = lp[1]
	}
	if ll > 2 {
		msg := fmt.Sprintf("Location ['%v'] can only have one colon (to denote version.)", locstr)
		err := errors.New(msg)
		return err
	}
	l.raw = locstr
	l.host = la[0]
	l.group = la[1]
	l.name = la[2]
	if l.name == constant.EmptyString {
		msg := fmt.Sprintf("Location ['%v'] cannot have an empty name.", locstr)
		err := errors.New(msg)
		return err
	}
	if l.style == constant.IntegerVersionStyle {
		iv, err := l.parseIntegerVersion(la[3])
		if err != nil {
			return err
		}
		l.version = strconv.Itoa(int(iv))
		l.integer = iv

	} else if l.style == constant.DottedVersionStyle {
		dv, err := l.parseDottedVersion(la[3])
		if err != nil {
			return err
		}
		l.version = dv.GetVersion()
	}
	l.parsed = true
	return nil
}

func (l *Location) parseIntegerVersion(verstr string) (byte, error) {
	iv, err := strconv.Atoi(verstr)
	if err != nil {
		msg := "Version ['%v'] in locator ['%v'] is not a valid integer: %v"
		msg = fmt.Sprintf(msg, verstr, l.raw, err)
		err = errors.New(msg)
		return 0, err
	}
	return byte(iv), nil
}

func (l *Location) parseDottedVersion(verstr string) (*DottedVersion, error) {
	dv := NewDottedVersion()
	err := dv.Parse(verstr)
	if err != nil {
		msg := "Invalid value ['%v'] in locator ['%v'] for dotted-style version: %v"
		msg = fmt.Sprintf(msg, verstr, l.raw, err)
		err = errors.New(msg)
		return nil, err
	}
	return dv, nil
}
