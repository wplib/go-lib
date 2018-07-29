package image

import (
	"io/ioutil"
	"github.com/wplib/go-lib/docker"

	"encoding/json"
)
func GetImageTags(token, image string) (docker.TagList,error) {
	p:= "/"+ image +"/tags/list"
	au:= docker.BuildApiUrl(docker.RegistryDomain,"2",p,"")
	client,req,err := docker.NewHttpGetRequest(au)
	if err != nil{
		return docker.TagList{},err
	}
	req.Header.Add("Authorization","Bearer "+token)
	resp, err := client.Do(req)
	if err != nil{
		return docker.TagList{},err
	}
	tags,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	tlr := docker.TagListResponse{}
	err = json.Unmarshal(tags,&tlr)
	return tlr.Tags,err

}


