package kitsuclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (k *Kitsu) GetUser(id string) (u User, err error) {
	rsp, err := http.Get(fmt.Sprintf("%s/users/%s", k.baseURL, id))
	if err != nil {
		return
	}

	var ur UserRequest
	err = json.NewDecoder(rsp.Body).Decode(&ur)
	if err != nil {
		return
	}

	return ur.Data, nil
}

//https://kitsu.io/api/edge/users?filter[name]=...
func (k *Kitsu) GetUserByName(name string) (u User, err error) {
	rsp, err := http.Get(fmt.Sprintf("%s/users?filter[name]=%s", k.baseURL, name))
	if err != nil {
		return
	}

	var ur UserFilterRequest
	err = json.NewDecoder(rsp.Body).Decode(&ur)
	if err != nil {
		return
	}

	return ur.Data[0], nil
}
