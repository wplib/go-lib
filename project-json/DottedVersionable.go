package project_json

import (
	"strings"
	"strconv"
	"fmt"
	"errors"
)

type DottedVersionable struct {
	parsed string
	defaults [3]byte
	major  byte
	minor  byte
	patch  byte
}

func NewDottedVersionable() *DottedVersionable {
	dv := &DottedVersionable{}
	dv.SetDefaults(0,0,0)
	dv.Reset()
	return dv
}

func (dv *DottedVersionable) SetDefaults(mj,mn,p byte) {
	dv.defaults = [3]byte{mj, mn, p}
}

func (dv *DottedVersionable) Reset() {
	dv.major = dv.defaults[0]
	dv.minor = dv.defaults[1]
	dv.patch = dv.defaults[2]
	dv.parsed = dv.GetVersion()
}

func (dv *DottedVersionable) Parse(sv string) error {
	va := dv.defaults
	vp := strings.Split(sv, ".")
	for i := 0; i <= 2; i++ {
		if (len(vp) < i) {
			continue
		}
		vn, err := strconv.Atoi(vp[i])
		if err != nil {
			return err
		}
		va[i] = byte(vn)
	}
	dv.parsed = sv
	for i:=0; i<2; i++ {
		vseg,err := strconv.Atoi(strconv.Itoa(int(va[i])))
		if err != nil {
			msg := fmt.Sprintf("Version segment ['%v'] in version ['%v'] is not valid: %v", vseg, va[i], err)
			err = errors.New(msg)
			return err
		}
		va[i] = byte(vseg)
	}
	dv.major = va[0]
	dv.minor = va[1]
	dv.patch = va[2]
	return nil
}

func (dv *DottedVersionable) GetParsed() string {
	return dv.parsed
}

func (dv *DottedVersionable) GetVersion() string {
	return dv.GetMajorMinor() + "." + strconv.Itoa(int(dv.patch))
}

func (dv *DottedVersionable) GetMajorMinor() string {
	return strconv.Itoa(int(dv.major)) + "." + strconv.Itoa(int(dv.minor))
}
