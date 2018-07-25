package project_json

/**
 * @todo Pull these from defaults somehow
 * These constants is only temporary
 */
const (
	_DefaultCTypeHost = "wplib.org"
	_DefaultCTypeStack = "wordpress"
	_DefaultCTypeName = ""
	_DefaultCTypeVersion = "1"
)

type ComponentType struct {
	ctype   string
	Locator
}

func NewComponentType() *ComponentType {
	l := NewLocatable(IntegerVersionStyle)
	l.SetDefaults(
		_DefaultCTypeHost,
		_DefaultCTypeStack,
		_DefaultCTypeName,
		_DefaultCTypeVersion,
	)
	return &ComponentType{
		Locator: l,
	}
}

func (ct *ComponentType) GetStack() string {
	return ct.GetGroup()
}

func (ct *ComponentType) GetVersion() int {
	return ct.GetIntegerVersion()
}

func (ct *ComponentType) GetStringVersion() string {
	return ct.Locator.GetVersion()
}

func (ct *ComponentType) GetType() string {
	return ct.GetHost() +
		"/" + ct.GetStack() +
		"/" + ct.GetName() +
		":" + ct.GetStringVersion()
}

