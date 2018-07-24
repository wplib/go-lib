package project_json

import (
	"strings"
	"strconv"
)

type ComponentVersion struct {
	Version string
	Major   byte
	Minor   byte
	Patch   byte
}

func NewComponentVersion() *ComponentVersion {
	return &ComponentVersion{
		Version: _DefaultCRefVersion,
		Major:   0,
		Minor:   0,
		Patch:   0,
	}
}

/**
 * @todo Debug this!
 */
func (cv *ComponentVersion) Parse(sv string) error {
	va := []byte{0, 0, 0}
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
	cv.Version = sv
	cv.Major = va[0]
	cv.Minor = va[1]
	cv.Patch = va[2]
	return nil
}

func (cv *ComponentVersion) FullVersion() string {
	return strconv.Itoa(int(cv.Major)) + "." +
		strconv.Itoa(int(cv.Minor)) + "." +
		strconv.Itoa(int(cv.Patch))
}
