package rmanager

import (
	"sync"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/helpers/request"
)

// requestContainer - Defines the single request struct, uses RemoteData to get data and a channel to write response or error if occured
type requestContainer struct {
	rd     request.RemoteData
	result chan interface{}
	err    chan error
}

// getData - Tries to get data from the remote resource
func (r requestContainer) getData() {

	content, err := request.GetRemoteData(r.rd)
	if err != nil {
		r.err <- err
	}
	r.result <- content

	close(r.result)
	close(r.err)
}

// RequestManager - Struct dedicated to managing a request queue,
// channel "next" is used to make the worker understand that he is ready to make a new request,
type RequestManager struct {
	requestQueue []requestContainer
	running      bool
	next         chan bool
	mu           sync.Mutex
}

var (
	requestManager     *RequestManager
	onceRequestManager sync.Once
)

// newRequestContainer - Returns an instance of requestContainer, accepts type that implements RemoteData
func newRequestContainer(r request.RemoteData) requestContainer {

	rc := *new(requestContainer)
	rc.rd = r
	rc.result = make(chan interface{}, 1)
	rc.err = make(chan error, 1)

	return rc
}

// GetRequestManager - Returns the once instance of RequestManager
func GetRequestManager() *RequestManager {

	onceRequestManager.Do(func() {
		requestManager = new(RequestManager)
	})
	return requestManager
}

// popFromQueue - Removes the first element from the queue
func (rm *RequestManager) popFromQueue() {

	rm.requestQueue = rm.requestQueue[1:]
	if len(rm.requestQueue) > 0 {
		rm.next <- true
	}
}

// work - It takes care of receiving requests and deleting the queue, loops Through "next" channel
func (rm *RequestManager) work() {

	rm.next = make(chan bool, 1)

	go func() {

		for range rm.next {

			rm.mu.Lock()
			rm.requestQueue[0].getData()
			rm.popFromQueue()
			rm.mu.Unlock()
		}
	}()
}

// AddRequest - Appends a request to the queue, if the worker is not running starts it and return the two channels of requestContainer to read the result or error
func (rm *RequestManager) AddRequest(r request.RemoteData) (<-chan interface{}, <-chan error) {

	rc := newRequestContainer(r)

	rm.mu.Lock()
	rm.requestQueue = append(rm.requestQueue, rc)
	rm.mu.Unlock()

	go func() {

		rm.mu.Lock()
		if !rm.running {
			rm.StartService()
		}
		rm.mu.Unlock()
	}()

	return rc.result, rc.err
}

// StartService - Starts the worker
func (rm *RequestManager) StartService() {

	if !rm.running {
		rm.running = true
		rm.work()
		rm.next <- true
	}
}
