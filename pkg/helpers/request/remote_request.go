package request

import (
	"encoding/json"
	"io"
	"net/http"
)

// RemoteData -	Interface to generalize a remote request
type RemoteData interface {
	EncodeQueryString(req *http.Request)
	GetBody() io.Reader
	GetMethod() string
	GetURL() string
}

// GetRemoteData - Exec the request
func GetRemoteData(r RemoteData) (interface{}, error) {

	req, err := http.NewRequest(r.GetMethod(), r.GetURL(), r.GetBody())
	if err != nil {
		return nil, err
	}

	r.EncodeQueryString(req)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var content interface{}

	if err := json.NewDecoder(res.Body).Decode(&content); err != nil {
		return nil, err
	}

	return content, nil
}
