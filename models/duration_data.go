package models

import (
	"errors"
	"sync"
	"time"
)

// DurationData - Struct per immagazzinare i dati raccolti con il suo relativo tempo di scadenza dopo il quale è obbligato a ricevere nuovi dati
type DurationData struct {
	Content   interface{}
	ExpiredAt time.Time
}

var data *DurationData
var onceData sync.Once

// GetData - Restituisce l'istanza di data
func GetData() *DurationData {
	onceData.Do(func() {
		data = new(DurationData)
	})
	return data
}

// GetContent - Restituisce i dati recuperati nel caso siano presenti e non siano scaduti altrimenti errore
func (w *DurationData) GetContent() (interface{}, error) {

	if w.ExpiredAt.IsZero() || w.Content == nil {
		return nil, errors.New("Dati mancanti")
	}

	diff := w.ExpiredAt.Sub(time.Now())
	if diff.Seconds() <= 0 {
		return nil, errors.New("Data scaduta")
	}
	return w.Content, nil
}

// SetContent - Imposta dei nuovi dati e aggiorando il tempo di scadenza solo se i precedenti non sono più validi, altrimenti non fa niente
func (w *DurationData) SetContent(content interface{}, minutsInterval int) {

	if w.ExpiredAt.IsZero() {
		w.Content = content
		w.ExpiredAt = time.Now().Add(time.Minute * time.Duration(minutsInterval))
		return
	}

	if diff := w.ExpiredAt.Sub(time.Now()); diff.Seconds() > 0 {
		return
	}

	w.Content = content
	w.ExpiredAt = time.Now().Add(time.Minute * time.Duration(minutsInterval))
}
