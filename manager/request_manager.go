package manager

import (
	"Go-Starter-Project/helpers/request"
	"sync"
)

// RequestManager -
type RequestManager struct {
	requestQueue []request.RemoteData
	running      bool
	next         chan bool
	result       chan bool
	mu           sync.Mutex
}

var requestManager *RequestManager
var onceRequestManager sync.Once

// GetRequestManager -
func GetRequestManager() *RequestManager {
	onceRequestManager.Do(func() {
		requestManager = new(RequestManager)
	})
	return requestManager
}

// popFromQueue - Si occupa di rimuovere il primo elemento dalla coda
func (rm *RequestManager) popFromQueue() {
	rm.requestQueue = rm.requestQueue[1:]
}

// work -
func (rm *RequestManager) work() {


	
}

// AddRequest - Si occupa di aggiungere una richiesta alla coda
func (rm *RequestManager) AddRequest(r request.RemoteData) {
	rm.requestQueue = append(rm.requestQueue, r)
}

// StartService -
func (rm *RequestManager) StartService() {
	rm.mu.Lock()
	if !rm.running {
		rm.running = true
	}
	rm.mu.Unlock()
}
