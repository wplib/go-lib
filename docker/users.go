package docker

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (a Auth) GetToken() (l Login) {
	uri := "/users/login/"
	bff := new(bytes.Buffer)
	json.NewEncoder(bff).Encode(a)
	req, err := http.NewRequest("POST", URL+uri, bff)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&l)
	return
}