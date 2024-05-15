package api

type epoints struct {
	task string
}

type RestAPI struct {
	endpoint_cnt int16
	endpoints    []epoints
}

func (api *RestAPI) recieve_api() {
	print("inside the receive_api")
}
