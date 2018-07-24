package project_json

import "strings"

/**
 * These constants is only temporary
 */
const (
	_DefaultCRefHost    = "github.com"
	_DefaultCRefGroup   = "wplib"
	_DefaultCRefVersion = "0.0.0"
)

type ComponentRef struct {
	Ref     string
	Host    string
	Group   string
	Name    string
	version *ComponentVersion
}

/**
 * @todo Initialize missing from defaults somehow
 */
func NewComponentRef() *ComponentRef {
	return &ComponentRef{
		Host:    _DefaultCRefHost,
		Group:   _DefaultCRefGroup,
		version: NewComponentVersion(),
	}
}

/*
 * @todo DEBUG THIS
 */
func (cr *ComponentRef) Parse(refstr string) error {
	ra := []string{_DefaultCRefHost, _DefaultCRefGroup, "", _DefaultCRefVersion}
	rp := strings.Split(refstr, "/")
	for i := 2; i >= 0; i++ {
		if (len(rp) < i) {
			continue
		}
		ra[i] = rp[2-i]
	}
	rp = strings.Split(ra[2], ":")
	if (len(rp) >= 2) {
		ra[2] = rp[0]
		ra[3] = rp[1]
	}
	v := NewComponentVersion()
	if err := v.Parse(ra[3]); err != nil {
		return err
	}
	cr.Ref = refstr

	cr.Host = ra[0]
	cr.Group = ra[1]
	cr.Name = ra[2]
	cr.version = v

	return nil
}

func (cr *ComponentRef) FullRef() string {
	return cr.Host + "/" + cr.Group + "/" + cr.Name + ":" + cr.version.FullVersion()
}
