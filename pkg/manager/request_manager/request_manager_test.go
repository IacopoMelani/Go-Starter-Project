package rmanager

import (
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/request"
)

// RemoteDataTest - Example
type RemoteDataTest struct{}

// EncodeQueryString - Example
func (u RemoteDataTest) EncodeQueryString(req *http.Request) {}

// GetBody - Example
func (u RemoteDataTest) GetBody() io.Reader {
	return nil
}

// GetMethod - Example
func (u RemoteDataTest) GetMethod() string {
	return "GET"
}

// GetURL - Example
func (u RemoteDataTest) GetURL() string {
	return "https://randomuser.me/api/"
}

// HandlerData - Example
func (u RemoteDataTest) HandlerData() (interface{}, error) {
	content, err := request.GetRemoteData(u)
	return content, err
}

func TestRequestManager(t *testing.T) {

	rm := GetRequestManager()

	for i := 0; i < 10; i++ {

		go func() {

			u := RemoteDataTest{}

			res, err := rm.AddRequest(u)

			select {
			case <-res:
				t.Log("response")
			case err := <-err:
				t.Log(err.Error())
			}

		}()
	}
	time.Sleep(4 * time.Second)
}
