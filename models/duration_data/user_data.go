package durationmodel

import (
	"io"
	"net/http"
	"sync"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/request"
	durationdata "github.com/IacopoMelani/Go-Starter-Project/pkg/models/duration_data"

	"github.com/IacopoMelani/Go-Starter-Project/config"
)

// UserRemoteData - Define a struct that implements DDInterface
type UserRemoteData struct{}

var userData *durationdata.DurationData
var onceUser sync.Once

// GetUsersData - Returns the instance of DurationData relative to user duration data
func GetUsersData() *durationdata.DurationData {
	onceUser.Do(func() {
		userData = new(durationdata.DurationData)
		userData.SetDurationDataInterface(UserRemoteData{})
		userData.SetTimeToRefresh(config.GetInstance().UserTimeToRefresh)
		userData.Daemon()
	})
	return userData
}

// EncodeQueryString - Define the to-implement func to add extra data to request header
func (u UserRemoteData) EncodeQueryString(req *http.Request) {}

// GetBody - Returns the body of request
func (u UserRemoteData) GetBody() io.Reader {
	return nil
}

// GetMethod - Returns the method of request ex "POST"
func (u UserRemoteData) GetMethod() string {
	return "GET"
}

// GetURL - Returns the endpoint of the resource
func (u UserRemoteData) GetURL() string {
	return "https://randomuser.me/api/"
}

// HandlerData - Exec the method handler to retrive the remote data
func (u UserRemoteData) HandlerData() (interface{}, error) {
	content, err := request.GetRemoteData(u)
	return content, err
}
