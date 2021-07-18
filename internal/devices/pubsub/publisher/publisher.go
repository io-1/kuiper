package publisher

type Publisher interface {
	Publish(topic string, json []byte)
}
