package project_json

import "github.com/wplib/project-cli/component_location"

type ComponentLocation struct {
	*Location
}

/**
 * @todo Initialize missing from defaults somehow
 */
func NewComponentLocation() *ComponentLocation {
	l := NewLocation(DottedVersionStyle)
	l.SetDefaults(
		component_location.DefaultHost(),
		component_location.DefaultGroup(),
		component_location.DefaultName(),
		component_location.DefaultVersion(),
	)
	return &ComponentLocation{
		Location: l,
	}
}

