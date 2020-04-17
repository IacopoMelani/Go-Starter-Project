package durationdata

import (
	"sync"
	"time"
)

// DDInterface - Interface to force implement the data handler
type DDInterface interface {
	HandlerData() (interface{}, error)
}

// DurationData - Struct to store the collected data with its relative expiration time, it is necessary to define a handler function to be assigned to the DurationData instance, a time interval in seconds in which the handler is called and then start the daemon relating to the same instance.
type DurationData struct {
	mu          sync.Mutex
	ddi         DDInterface
	stopSignal  chan bool
	sleepSecond int
	content     interface{}
}

var registeredInitDurationData []func() *DurationData

// InitDurationData - It takes care of starting all instances of DurationData
func InitDurationData() {
	for _, f := range registeredInitDurationData {
		f()
	}
}

// RegisterInitDurationData - Records the functions that start their DurationData
func RegisterInitDurationData(f ...func() *DurationData) {
	registeredInitDurationData = append(registeredInitDurationData, f...)
}

// getDaemonData - It takes care of pre-fetching the data from the handler and if there have been no errors it replaces it with the new one
func (d *DurationData) getDaemonData() {
	content, err := d.ddi.HandlerData()
	if err == nil {
		d.SetSafeContent(content)
	}
}

// Daemon - It takes care of starting the daemon that updates the data, it can be killed by calling the StopDaemon() func
func (d *DurationData) Daemon() {

	d.stopSignal = make(chan bool)

	go func() {
		ticker := time.NewTicker(time.Second * time.Duration(d.sleepSecond))

		d.getDaemonData()

		for {
			select {
			case <-d.stopSignal:
				ticker.Stop()
				return
			case <-ticker.C:
				d.getDaemonData()

			}
		}
	}()
}

// GetSafeContent - It exclusively returns the duration data content
func (d *DurationData) GetSafeContent() interface{} {

	d.mu.Lock()
	content := d.content
	d.mu.Unlock()

	return content
}

// SetSafeContent - Set up your data securely
func (d *DurationData) SetSafeContent(content interface{}) {
	d.mu.Lock()
	d.content = content
	d.mu.Unlock()
}

// SetDurationDataInterface - Set the struct that implements DDInterface
func (d *DurationData) SetDurationDataInterface(ddi DDInterface) {
	d.ddi = ddi
}

// SetTimeToRefresh - Sets the value of the remote data refresh time
func (d *DurationData) SetTimeToRefresh(t int) {
	d.sleepSecond = t
}

// StopDaemon - It takes care of warning the demon to stop
func (d *DurationData) StopDaemon() {
	d.stopSignal <- true
	close(d.stopSignal)
}
