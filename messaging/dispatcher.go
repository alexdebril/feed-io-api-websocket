package messaging

type Dispatcher interface {
	Handle(item Item)
}
