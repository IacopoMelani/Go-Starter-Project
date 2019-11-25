package durationmodel

import (
	"io"
	"net/http"
	"sync"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/request"
	durationdata "github.com/IacopoMelani/Go-Starter-Project/pkg/models/duration_data"

	"github.com/IacopoMelani/Go-Starter-Project/config"
)

// UserRemoteData - Definisce una struct che implementa RemoteData
type UserRemoteData struct{}

var userData *durationdata.DurationData
var onceUser sync.Once

// GetUsersData - Restituisce l'istanza di DurantionData relativo agli utenti
func GetUsersData() *durationdata.DurationData {
	onceUser.Do(func() {
		userData = new(durationdata.DurationData)
		userData.SetDurationDataInterface(UserRemoteData{})
		userData.SetTimeToRefresh(config.GetInstance().UserTimeToRefresh)
		userData.Daemon()
	})
	return userData
}

// EncodeQueryString - Si occupa di aggiungere i paratri dell'header alla request
func (u UserRemoteData) EncodeQueryString(req *http.Request) {}

// GetBody - Restituisce il corpo della request
func (u UserRemoteData) GetBody() io.Reader {
	return nil
}

// GetMethod - Restituisce il metodo della richiesta remota
func (u UserRemoteData) GetMethod() string {
	return "GET"
}

// GetURL - Restituisce la url della richiesta remota
func (u UserRemoteData) GetURL() string {
	return "https://randomuser.me/api/"
}

// HandlerData - Si occupa di eseguire la funzione di handler per ricevere i dati
func (u UserRemoteData) HandlerData() (interface{}, error) {
	content, err := request.GetRemoteData(u)
	return content, err
}
