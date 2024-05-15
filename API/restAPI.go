package api

import (
	"net/http"
)

type epoints struct {
	task    string
	process func(http.ResponseWriter, *http.Request)
}

type RestAPI struct {
	endpoint_cnt int16
	endpoints    map[string]epoints
}

// optional
func (api *RestAPI) recieve_api() {
	print("inside the receive_api")
}
