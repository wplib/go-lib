package project_json

/**
 * These constants is only temporary
 */
const (
	_DefaultCRefHost    = "github.com"
	_DefaultCRefGroup   = "wplib"
	_DefaultCRefName    = ""
	_DefaultCRefVersion = "0.0.0"
)

type ComponentRef struct {
	Locator
}

/**
 * @todo Initialize missing from defaults somehow
 */
func NewComponentRef() *ComponentRef {
	l := NewLocatable(DottedVersionStyle)
	l.SetDefaults(
		_DefaultCRefHost,
		_DefaultCRefGroup,
		_DefaultCRefName,
		_DefaultCRefVersion,
	)
	return &ComponentRef{
		Locator: l,
	}
}

