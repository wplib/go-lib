package project_json


import (
	"testing"
)

const errorExpected= "~"

func newTestableIntegerVersionLocation() *Location {
	l := NewLocation(IntegerVersionStyle)
	l.SetDefaults( "wplib.org", "wordpress", EmptyString, "1" )
	return l
}

var integerVersionLocationTests = []struct {
	in  string
	out string
}{
	{"wplib.org/wordpress/webserver:1","wplib.org/wordpress/webserver:1"},
	{"wplib.org/lxmp/webserver","wplib.org/lxmp/webserver:1"},
	{"wplib.org/wordpress/webserver","wplib.org/wordpress/webserver:1"},
	{"wordpress.org/wordpress/webserver","wordpress.org/wordpress/webserver:1"},
	{"wordpress/webserver","wplib.org/wordpress/webserver:1"},
	{"webserver","wplib.org/wordpress/webserver:1"},
	{"webserver:2","wplib.org/wordpress/webserver:2"},
	{EmptyString,errorExpected},
	{"webserver:abc",errorExpected},
	{":1",errorExpected},
	{"webserver:1:0",errorExpected},
	{"world/wplib.org/wordpress/webserver:1",errorExpected},
}

func TestIntegerVersionLocation(t *testing.T) {
	for _, tt := range integerVersionLocationTests {
		l:= newTestableIntegerVersionLocation()
		t.Run(tt.in, func(t *testing.T) {
			err := l.Parse(tt.in)
			if tt.out == errorExpected {
				if err == nil {
					t.Errorf("wanted error %q, did not get", tt.in)
				}
			} else {
				ls := l.GetLocation()
				if ls != tt.out {
					t.Errorf("wanted %q, got %q", ls, tt.out)
				}
			}
		})
	}
}