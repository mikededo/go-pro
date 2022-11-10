package bus

type Payload[T any] interface{}

type Bus interface {
	// Publish publishes arguments to the given topic subscribers
	Publish(topic string, args ...interface{})
	// Unpublish unsubscribe all handlers from given topic
	Unpublish(topic string)
	// Subscribe subscribes to the given topic
	Subscribe(topic string, fn interface{}) error
	// Unsubscribe unsubscribe handler from the given topic
	Unsubscribe(topic string, fn interface{}) error
}
