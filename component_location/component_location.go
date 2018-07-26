package component_location

import "github.com/wplib/project-cli/constant"

func DefaultHost() string {
	return "github.com"
}
func DefaultGroup() string {
	return "wplib"
}
func DefaultName() string {
	return ""
}
func DefaultVersion() string {
	return constant.DefaultDottedVersion
}