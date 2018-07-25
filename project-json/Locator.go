package project_json

type Locator interface {
	SetDefaults(h,g,i,v string)
	GetDefaults() LocatableDefaults
	GetParsed() string
	GetHost() string
	GetGroup() string
	GetName() string
	GetVersion() string
	GetIntegerVersion() int
	GetDottedVersionable() *DottedVersionable
	GetDottedVersion() string
	GetLocator() string
	Parse(locstr string) error
}