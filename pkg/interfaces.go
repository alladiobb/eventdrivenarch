package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayLoad() interface {}
}

type EventHandlerInterface interface {
	Handle(event EventInterface)
}

type EventDispatcherInterface interface{
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventHandlerInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}