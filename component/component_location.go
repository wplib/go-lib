package component

import (
	"github.com/wplib/go-lib/location"
	"github.com/wplib/go-lib/constant"
	"github.com/wplib/go-lib/defaults"
)

type ComponentLocation struct {
	*location.Location
}

func NewComponentLocation() *ComponentLocation {
	l := location.NewLocation(constant.DottedVersionStyle)
	l.SetDefaults(
		defaults.DefaultLocationHost(),
		defaults.DefaultLocationGroup(),
		defaults.DefaultLocationName(),
		defaults.DefaultLocationVersion(),
	)
	return &ComponentLocation{
		Location: l,
	}
}

