package controllers

import (
	"github.com/IacopoMelani/Go-Starter-Project/config"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/subosito/gotenv"
)

func TestGetAllUser(t *testing.T) {

	gotenv.Load("../.env")

	config := config.GetInstance()

	client := http.Client{}

	port := config.AppPort

	req, err := http.NewRequest("GET", "http://localhost"+port+"/user/all", nil)
	if err != nil {
		t.Error("Errore: impossibile creare la richiesta")
	}

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Errore: impossiblie eseguire la richiesta url: %s, potrebbe essere necessario avviare il server", "http://localhost:"+port+"/user/all")
	}
	defer res.Body.Close()

	var response Response

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Error("Errore: impossibile leggere dalla response")
	}

	if response.Status != 0 || response.Success != true {
		t.Errorf("Errore: si è verificato un errore nella richiesta codice: %d  messaggio: %s", response.Status, response.Message)
	}
}
