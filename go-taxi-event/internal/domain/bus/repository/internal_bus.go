package repository

import (
	"fmt"
	"reflect"
	"sync"
)

type handler struct {
	callback reflect.Value
	queue    chan []reflect.Value
}

type topicHandlers map[string][]*handler

type InternalBus struct {
	mtx      sync.RWMutex
	buffSize int
	handlers topicHandlers
}

func newHandler(buffSize int, cb interface{}) *handler {
	return &handler{
		callback: reflect.ValueOf(cb),
		queue:    make(chan []reflect.Value, buffSize),
	}
}

func NewEventBus(buffSize int) *InternalBus {
	return &InternalBus{
		buffSize: buffSize,
		handlers: make(topicHandlers),
	}
}

func (b *InternalBus) Subscribe(topic string, cb interface{}) error {
	if err := isValidHandler(cb); err != nil {
		return err
	}

	h := newHandler(b.buffSize, cb)

	// init the go routine that will listed for events and
	// send them to the handler
	go func() {
		// receive the arguments from the chan
		for args := range h.queue {
			// send them to the listener
			h.callback.Call(args)
		}
	}()

	// add the handler to the handlers
	b.mtx.Lock()
	defer b.mtx.Unlock()

	b.handlers[topic] = append(b.handlers[topic], h)

	return nil
}

func (b *InternalBus) Unsubscribe(topic string, cb interface{}) error {
	if err := isValidHandler(cb); err != nil {
		return err
	}

	hs, ok := b.handlers[topic]
	if ok {
		return fmt.Errorf("%s topic does not exist", topic)
	}

	b.mtx.Lock()
	defer b.mtx.Unlock()

	rv := reflect.ValueOf(cb)
	for i, h := range hs {
		if h.callback == rv {
			// close the channel
			close(h.queue)

			// if only one listener, remove the topic
			if len(hs) == 1 {
				delete(b.handlers, topic)
			} else {
				hs = append(hs[:i], hs[i+1:]...)
			}
		}
	}

	return nil

}

func (b *InternalBus) Notify(topic string, payload ...interface{}) {
	gp := generatePayload(payload)

	// lock mutex for read
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	// send payload to listener
	if hs, ok := b.handlers[topic]; ok {
		for _, h := range hs {
			h.queue <- gp
		}
	}
}

func (b *InternalBus) Clear(topic string) {
	hs, ok := b.handlers[topic]
	if !ok {
		return
	}

	b.mtx.Lock()
	defer b.mtx.Unlock()

	// close all handlers
	for _, h := range hs {
		close(h.queue)
	}
	// delete the entry in the map
	delete(b.handlers, topic)
}

func isValidHandler(cb interface{}) error {
	if t := reflect.TypeOf(cb); t.Kind() != reflect.Func {
		return fmt.Errorf("%s is not reflect.Func", t)
	}

	return nil
}

func generatePayload(payload []interface{}) []reflect.Value {
	rp := make([]reflect.Value, 0)
	for _, p := range payload {
		rp = append(rp, reflect.ValueOf(p))
	}
	return rp
}
