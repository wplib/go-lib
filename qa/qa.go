package qa

import "testing"

type Any interface{}
	

const ErrExpected = "~"

type InOut interface {
	Input() string
	Output() string
}

type TestHarness struct {
	T *testing.T
	InOut
	Item Any
}
func NewTestHarness(t *testing.T,td InOut,i Any) *TestHarness {
	return &TestHarness{
		T:t,
		InOut:td,
		Item:i,
	}
}
func (th *TestHarness) Run(fn func()) {
	th.T.Run(th.Input(), func(t *testing.T) {
		fn()
	})
}


