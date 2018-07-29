package auth

import (
	"github.com/wplib/go-lib/docker"
	"io/ioutil"
	"net/url"
	"encoding/json"
)


func GetAuthToken(repo,scope string) (string,error){
	p:= "/token"
	q:= url.Values{
		"service":{"registry.docker.io"},
		"scope":{"repository:"+repo+":"+scope},
	}
	au:= docker.BuildApiUrl(docker.AuthDomain,"",p,q.Encode())
	client,req,err := docker.NewHttpGetRequest(au)
	if err != nil{
		return "",err
	}
	resp, err := client.Do(req)
	if err != nil{
		return "",err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return "",err
	}
	authToken := docker.AuthToken{}
	err = json.Unmarshal(body,&authToken)
	return authToken.Token,err
}


//func foo(){
//	var netTransport = &http.Transport{
//		Dial: (&net.Dialer{
//			Timeout: 5 * time.Second,
//		}).Dial,
//		TLSHandshakeTimeout: 5 * time.Second,
//	}
//	var netClient = &http.Client{
//		Timeout: time.Second * 10,
//		Transport: netTransport,
//	}
//	response, _ := netClient.Get(url)
//}