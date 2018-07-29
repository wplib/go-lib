package docker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetBuildHistory(a Auth, user, repo string) (bh BuildHistory) {
	uri := fmt.Sprintf("/repositories/%s/%s/buildhistory/", user, repo)
	req, err := http.NewRequest("GET", URL+uri, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", a.GetToken()))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&bh)
	return
}