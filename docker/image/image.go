package image

import (
	"github.com/wplib/go-lib/docker/api"
	"github.com/wplib/go-lib/docker/tag"
	"github.com/wplib/go-lib/docker"
	"encoding/json"
	"io/ioutil"
)

func GetRemoteImageTagTree(token, image string) (tag.TagTree,error) {
	tags,err:= GetRemoteImageTags(token,image)
	if err != nil{
		return tag.TagTree{},err
	}
	return tag.NewTagTree(image,tags)
}

func GetRemoteImageTags(token, image string) (tag.TagList,error) {

	au:= api.BuildUrl(
		docker.RegistryDomain,
		docker.HubApiVersion,
		"/"+ image +"/tags/list",
		docker.EmptyQueryList,
	)

	client,req,err := api.NewHttpGetRequest(au)
	if err != nil{
		return tag.TagList{},err
	}
	req.Header.Add("Authorization","Bearer "+token)
	resp, err := client.Do(req)
	if err != nil{
		return tag.TagList{},err
	}
	tags,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	tlr := tag.TagListResponse{}
	err = json.Unmarshal(tags,&tlr)
	return tlr.Tags,err

}


