package rmanager

import (
	"Go-Starter-Project/helpers/request"
	"sync"
)

// requestContainer - Struct che defisce una singola richiesta, definisce l'istanza di RemoteData da cui prelevare i dati e un channel su cui scrivere il risultato
type requestContainer struct {
	rd     request.RemoteData
	result chan interface{}
	err    chan error
}

// getData - Si occupa di contattare la risorsa remota tramite RemoteData e in caso di successo, scrivere il risultato
func (r requestContainer) getData() {

	content, err := request.GetRemoteData(r.rd)
	if err != nil {
		r.err <- err
	}
	r.result <- content

	close(r.result)
	close(r.err)
}

// RequestManager - Struct dedicata alla gestione di una coda di richieste,
// il channel next serve per far capire al worker che è pronto per eseguire una nuova richiesta,
// il channeò stopSignal viene usato per interrompere il worker, il suo valore determina se le restanti richieste in coda verrano cancellate o meno
type RequestManager struct {
	requestQueue []requestContainer
	running      bool
	next         chan bool
	mu           sync.Mutex
}

var requestManager *RequestManager
var onceRequestManager sync.Once

// newrequestContainer - Restituisce un'istanza di requestContainer prendendo un'istanza che implementa RemoteData
func newrequestContainer(r request.RemoteData) requestContainer {

	rc := *new(requestContainer)
	rc.rd = r
	rc.result = make(chan interface{}, 1)
	rc.err = make(chan error, 1)

	return rc
}

// GetRequestManager - Restituisce l'istanza di RequestManager
func GetRequestManager() *RequestManager {

	onceRequestManager.Do(func() {
		requestManager = new(RequestManager)
	})
	return requestManager
}

// popFromQueue - Si occupa di rimuovere il primo elemento dalla coda
func (rm *RequestManager) popFromQueue() {

	rm.requestQueue = rm.requestQueue[1:]
	if len(rm.requestQueue) > 0 {
		rm.next <- true
	}
}

// work - Si occupa prelevare le richieste eseguirle e successivamente eliminarle dalla coda
func (rm *RequestManager) work() {

	rm.next = make(chan bool, 1)

	go func() {

		for {
			select {
			case <-rm.next:
				rm.requestQueue[0].getData()
				rm.popFromQueue()
			}
		}
	}()
}

// AddRequest - Si occupa di aggiungere una richiesta alla coda
func (rm *RequestManager) AddRequest(r request.RemoteData) (<-chan interface{}, <-chan error) {

	rc := newrequestContainer(r)

	rm.requestQueue = append(rm.requestQueue, rc)

	go func() {

		if !rm.running {
			rm.StartService()
		}
	}()

	return rc.result, rc.err
}

// StartService - Si occupa di avviare il servizio di code
func (rm *RequestManager) StartService() {

	rm.mu.Lock()
	if !rm.running {
		rm.running = true
		rm.work()
		rm.next <- true
	}
	rm.mu.Unlock()
}
