package sunday

type FrameworkContext interface {
	Bind(interface{}) error
	JSON(int, interface{})
}
