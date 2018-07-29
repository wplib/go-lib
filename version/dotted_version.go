package version

import (
	"strings"
	"strconv"
	"fmt"
	"errors"
	"github.com/wplib/go-lib/constant"
)
type DottedVersion struct {
	raw      string
	defaults [3]byte
	major    byte
	minor    byte
	patch    byte
}

func NewDottedVersion() *DottedVersion {
	dv := &DottedVersion{}
	dv.SetDefaults(0,0,0)
	dv.Reset()
	return dv
}

func (dv *DottedVersion) SetDefaults(mj,mn,p byte) {
	dv.defaults = [3]byte{mj, mn, p}
}

func (dv *DottedVersion) Reset() {
	dv.major = dv.defaults[0]
	dv.minor = dv.defaults[1]
	dv.patch = dv.defaults[2]
	dv.raw = dv.GetVersion()
}

func (dv *DottedVersion) Parse(verstr string) error {
	va := dv.defaults
	vp := strings.Split(verstr+"."+constant.DefaultDottedVersion, ".")
	for i := 0; i <= 2; i++ {
		if (len(vp) <= i) {
			break
		}
		vn, err := strconv.Atoi(vp[i])
		if err != nil {
			return err
		}
		va[i] = byte(vn)
	}
	dv.raw = verstr
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

func (dv *DottedVersion) GetRaw() string {
	return dv.raw
}

func (dv *DottedVersion) GetRawVersion() string {
	return dv.raw
}

func (dv *DottedVersion) GetVersion() string {
	return dv.GetMajorMinor() + "." + strconv.Itoa(int(dv.patch))
}

func (dv *DottedVersion) GetMajorMinor() string {
	return strconv.Itoa(int(dv.major)) + "." + strconv.Itoa(int(dv.minor))
}
