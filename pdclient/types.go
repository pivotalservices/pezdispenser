package pdclient

import "net/http"

type (
	//PDClient - dispenser client object
	PDClient struct {
		APIKey string
		client clientDoer
		URL    string
	}
	clientDoer interface {
		Do(req *http.Request) (resp *http.Response, err error)
	}
)
