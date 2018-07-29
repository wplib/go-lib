package test


import (
	"testing"
	"github.com/wplib/go-lib/qa"
	"github.com/wplib/go-lib/constant"
	"github.com/wplib/go-lib/location"
)

type testDefaults struct {
	host string
	group string
	name string
	ver string
	style byte
}
var integer = testDefaults{ "wplib.org", "wordpress", constant.EmptyString, "1", constant.IntegerVersionStyle }
var dotted = testDefaults{ "github.com", "wplib", constant.EmptyString, "0.0.0", constant.DottedVersionStyle}

type locationData struct {
	defaults testDefaults
	in  string
	out string
	host string
	group string
	name string
}
func (ld locationData) Input() string {
	return ld.in
}
func (ld locationData) Output() string {
	return ld.out
}
var locationTests = []locationData{
	{integer,constant.EmptyString, constant.ErrExpected,"","",""},
	{integer,"webserver:abc", constant.ErrExpected,"","",""},
	{integer,":1", constant.ErrExpected,"","",""},
	{integer,"webserver:1:0", constant.ErrExpected,"","",""},
	{integer,"world/wplib.org/wordpress/webserver:1", constant.ErrExpected,"","",""},
	{integer,"wplib.org/lxmp/webserver","wplib.org/lxmp/webserver:1","wplib.org","lxmp","webserver"},
	{integer,"wplib.org/wordpress/webserver:1","wplib.org/wordpress/webserver:1","wplib.org","wordpress","webserver"},
	{integer,"wplib.org/wordpress/webserver","wplib.org/wordpress/webserver:1","wplib.org","wordpress","webserver"},
	{integer,"wordpress.org/jetpack/webserver","wordpress.org/jetpack/webserver:1","wordpress.org","jetpack","webserver"},
	{integer,"wordpress/webserver","wplib.org/wordpress/webserver:1","wplib.org","wordpress","webserver"},
	{integer,"webserver","wplib.org/wordpress/webserver:1","wplib.org","wordpress","webserver"},
	{integer,"webserver:2","wplib.org/wordpress/webserver:2","wplib.org","wordpress","webserver"},
	{dotted,constant.EmptyString, constant.ErrExpected,"","",""},
	{dotted,"nginx:abc", constant.ErrExpected,"","",""},
	{dotted,":1.14.0", constant.ErrExpected,"","",""},
	{dotted,"nginx:1.14.0:0", constant.ErrExpected,"","",""},
	{dotted,"world/github.com/wplib/nginx:1.14.0", constant.ErrExpected,"","",""},
	{dotted,"nginx:2", "github.com/wplib/nginx:2.0.0","github.com","wplib","nginx"},
	{dotted,"github.com/wplib/nginx:1.14.0", "github.com/wplib/nginx:1.14.0","github.com","wplib","nginx"},
	{dotted,"github.com/lxmp/nginx", "github.com/lxmp/nginx:0.0.0","github.com","lxmp","nginx"},
	{dotted,"github.com/wplib/nginx", "github.com/wplib/nginx:0.0.0","github.com","wplib","nginx"},
	{dotted,"git.wplib.org/wplib/nginx", "git.wplib.org/wplib/nginx:0.0.0","git.wplib.org","wplib","nginx"},
	{dotted,"wplib/nginx", "github.com/wplib/nginx:0.0.0","github.com","wplib","nginx"},
	{dotted,"nginx", "github.com/wplib/nginx:0.0.0","github.com","wplib","nginx"},
}

func TestLocation(t *testing.T) {
	for _, ld := range locationTests {
		d := ld.defaults
		l := location.NewLocation(d.style)
		l.SetDefaults( d.host, d.group, d.name, d.ver )
		th:= qa.NewTestHarness(t,ld,l)
		th.Run(func(){
			if parseTest(th) == nil {
				nameTest(th)
				groupTest(th)
				hostTest(th)
				locationTest(th)
			}
		})
	}
}

func getLocation(th *qa.TestHarness) *location.Location {
	return th.Item.(*location.Location)
}

func parseTest(th *qa.TestHarness) error {
	err := getLocation(th).Parse(th.Input())
	if th.Output() == constant.ErrExpected {
		if err == nil {
			th.T.Errorf("wanted error %q, did not get", th.Input())
		}
	}
	return err
}

func locationTest(th *qa.TestHarness) {
	l:= th.Item.(*location.Location)
	ls:= l.GetLocation()
	if ls != th.Output() {
		th.T.Errorf("wanted %q, got %q", th.Output(), ls)
	}
}

func hostTest(th *qa.TestHarness) {
	l := th.Item.(*location.Location)
	d := th.InOut.(locationData)
	h := l.GetHost()
	if h != d.host {
		th.T.Errorf("wanted %q, got %q", d.host, h)
	}
}

func groupTest(th *qa.TestHarness) {
	l:= th.Item.(*location.Location)
	d:= th.InOut.(locationData)
	g:= l.GetGroup()
	if g != d.group {
		th.T.Errorf("wanted %q, got %q", d.group, g)
	}
}
func nameTest(th *qa.TestHarness) {
	l:= th.Item.(*location.Location)
	d:= th.InOut.(locationData)
	g:= l.GetName()
	if g != d.name {
		th.T.Errorf("wanted %q, got %q", d.name, g)
	}
}

