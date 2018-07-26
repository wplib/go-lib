package project_json


import (
	"testing"
	"github.com/wplib/project-cli/constant"
)

const errorExpected= "~"

func newTestableIntegerVersionLocation() *Location {
	l := NewLocation(constant.IntegerVersionStyle)
	l.SetDefaults( "wplib.org", "wordpress", constant.EmptyString, "1" )
	return l
}

var integerVersionLocationTests = []struct {
	in  string
	out string
}{
	{"wplib.org/lxmp/webserver","wplib.org/lxmp/webserver:1"},
	{"wplib.org/wordpress/webserver:1","wplib.org/wordpress/webserver:1"},
	{"wplib.org/wordpress/webserver","wplib.org/wordpress/webserver:1"},
	{"wordpress.org/wordpress/webserver","wordpress.org/wordpress/webserver:1"},
	{"wordpress/webserver","wplib.org/wordpress/webserver:1"},
	{"webserver","wplib.org/wordpress/webserver:1"},
	{"webserver:2","wplib.org/wordpress/webserver:2"},
	{constant.EmptyString,errorExpected},
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
func newTestableDottedVersionLocation() *Location {
	l := NewLocation(constant.DottedVersionStyle)
	l.SetDefaults( "github.com", "wplib", constant.EmptyString, "0.0.0" )
	return l
}

var dottedVersionLocationTests = []struct {
	in  string
	out string
}{
	{"nginx:2","github.com/wplib/nginx:2.0.0"},
	{"github.com/wplib/nginx:1.14.0","github.com/wplib/nginx:1.14.0"},
	{"github.com/lxmp/nginx","github.com/lxmp/nginx:0.0.0"},
	{"github.com/wplib/nginx","github.com/wplib/nginx:0.0.0"},
	{"wplib.org/wplib/nginx","wplib.org/wplib/nginx:0.0.0"},
	{"wplib/nginx","github.com/wplib/nginx:0.0.0"},
	{"nginx","github.com/wplib/nginx:0.0.0"},
	{constant.EmptyString,errorExpected},
	{"nginx:abc",errorExpected},
	{":1.14.0",errorExpected},
	{"nginx:1.14.0:0",errorExpected},
	{"world/github.com/wplib/nginx:1.14.0",errorExpected},
}

func TestDottedVersionLocation(t *testing.T) {
	for _, tt := range dottedVersionLocationTests {
		l:= newTestableDottedVersionLocation()
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

