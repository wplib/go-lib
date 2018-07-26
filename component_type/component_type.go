package component_type

import "github.com/wplib/project-cli/constant"

func DefaultHost() string {
	return "wplib.org"
}
func DefaultStack() string {
	return "wordpress"
}
func DefaultName() string {
	return ""
}
func DefaultVersion() string {
	return constant.DefaultIntegerVersion
}