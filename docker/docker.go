package docker

import (
	"net/http"
	"context"

	"strings"
)
const (
	AuthDomain ="auth.docker.io"
	RegistryDomain="registry.hub.docker.com"
	HubDomain ="hub.docker.com"
)

const UserAgent = "wplib-box-cli-0.1"

type HttpHeaders map[string]string

type Tree interface {
	Tag() string
	Children() *map[string]TagTree
}
type TagList []string

type TagTree struct {
	tag string
	children map[string]TagTree
}
func (tt TagTree) Tag() string {
	return tt.tag
}
func (tt TagTree) Children() *map[string]TagTree {
	return &tt.children
}
func (tt TagTree) GetChild(tag string) *TagTree {
	tc:= tt.children[tag]
	return &tc
}
func NewEmptyTagTree(tag string) *TagTree {
	return &TagTree{
		tag:tag,
		children:make(map[string]TagTree),
	}
}
func NewTagTree(list TagList) (TagTree,error) {
	tt:= NewEmptyTagTree("<root>")
	for _,t := range list {
		if ( t=="latest") {
			continue
		}
		ttt:= tt
		nums := strings.Split(t,".")
		var ok bool
		for _,n := range nums {
			c:= *ttt.Children()
			if _,ok = c[n]; ! ok {
				ttt.children[n] = *NewEmptyTagTree(n);
			}
			ttt= ttt.GetChild(n)
		}
	}
	return *tt,nil
}

func PrintTree(t Tree,dd ...byte) {
	var d int
	if len(dd)==0 {
		d= 0
	} else {
		d= int(dd[0])
	}
	for _,tag := range *t.Children() {
		println(strings.Repeat("-",d*3)+" "+tag.tag)
		PrintTree(tag,byte(d+1))
	}
}


type TagListResponse struct {
	Name string `json:"name"`
	Tags TagList `json:"tags"`
}

type AuthToken struct {
	Token string `json:"token"`
	AccessToken string  `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
	IssuedAt string `json:"issued_at"`
}

func GetContext() context.Context {
	return context.Background()
}

func CharAt(s string, b byte) byte {
	sa:=[]rune(s)
	return byte(sa[b])
}

func NewHttpGetRequest(url string) (*http.Client, *http.Request, error) {
	return NewHttpRequest("GET",url)
}

func NewHttpRequest(method,url string,) (*http.Client, *http.Request, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err == nil{
		req.Header.Add("User-Agent",UserAgent)
		req.Header.Add("Content-Type","application/json")
	}
	return client,req,err
}

func BuildApiUrl(host, version, path, query string) string {
	scheme := "https"
	if version != "" {
		version = "/v"+version
	}
	if CharAt(path,0) != '/' {
		path = "/"+ path
	}
	if query != "" {
		query = "?"+ query
	}
	return scheme+"://"+host+version+path+query
}

