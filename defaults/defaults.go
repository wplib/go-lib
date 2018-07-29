package defaults

import "github.com/wplib/go-lib/constant"

func DefaultLocationHost() string {
	return "github.com"
}
func DefaultLocationGroup() string {
	return "wplib"
}
func DefaultLocationName() string {
	return ""
}
func DefaultLocationVersion() string {
	return constant.DefaultDottedVersion
}

func DefaultTypeHost() string {
	return "wplib.org"
}
func DefaultTypeStack() string {
	return "wordpress"
}
func DefaultTypeName() string {
	return ""
}
func DefaultTypeVersion() string {
	return constant.DefaultIntegerVersion
}