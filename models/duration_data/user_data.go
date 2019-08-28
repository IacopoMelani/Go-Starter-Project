package durationdata

import (
	"github.com/Go-Starter-Project/config"
	"github.com/Go-Starter-Project/helpers/request"
	"io"
	"net/http"
	"sync"
)

// UserRemoteData - Definisce una struct che implementa RemoteData
type UserRemoteData struct{}

var userData *DurationData
var onceUser sync.Once

// GetUsersData - Restituisce l'istanza di DurantionData relativo agli utenti
func GetUsersData() *DurationData {
	onceUser.Do(func() {
		userData = new(DurationData)
		userData.ddi = UserRemoteData{}
		userData.sleepSecond = config.GetInstance().UserTimeToRefresh
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
