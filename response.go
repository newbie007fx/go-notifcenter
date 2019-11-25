package notif

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Error error
	resp  *http.Response
}

func (r Response) GetAsMap() (*map[string]interface{}, error) {
	if r.Error != nil {
		return nil, r.Error
	}
	res := map[string]interface{}{}
	err := json.NewDecoder(r.resp.Body).Decode(&res)
	return &res, err
}

func (r Response) GetAsStruct() (*NotifResponse, error) {
	if r.Error != nil {
		return nil, r.Error
	}
	res := NotifResponse{}
	err := json.NewDecoder(r.resp.Body).Decode(&res)
	return &res, err
}
