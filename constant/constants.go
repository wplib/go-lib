package constant

const EmptyString = ""

const (
	DefaultIntegerVersion = "1"
	DefaultDottedVersion = "0.0.0"
)

const (
	IntegerVersionStyle = iota
	DottedVersionStyle
)

var versionStyleMap = map[int]string {
	IntegerVersionStyle: DefaultIntegerVersion,
	DottedVersionStyle: DefaultDottedVersion,
}

func GetVersionStyleDefault(style byte) string {
	return versionStyleMap[int(style)]
}


const (
	//StackComponent      = iota
	ServiceComponent      = iota
	//ExecutableComponent
	//ScriptComponent
	//SourceComponent
	//DataComponent
	//MediaComponent
)
