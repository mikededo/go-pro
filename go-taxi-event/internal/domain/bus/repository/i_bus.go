package repository

type Bus interface {
	// Nofity notifies an event to all listeners for the topic
	Notify(topic string, payload ...interface{})
	// Clear clears all handles from a topic
	Clear(topic string)
	// Subscribe subscribes a handler to a topic
	Subscribe(topic string, cb interface{}) error
	// Unsubscribe unsubscribes a handler from a topic
	Unsubscribe(topic string, cb interface{}) error
}
