package events

import ("errors"
		"sync"
		)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct{
	handlers map[string][]EventHandlerInterface
    mu       sync.RWMutex
}

func NewEventDispatcher() *EventDispatcher{
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error{
	
	ed.mu.Lock()
    defer ed.mu.Unlock()
	
	if _, ok := ed.handlers[eventName]; ok{
		for _, h := range ed.handlers[eventName]{
			if h == handler{
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Clear(){
	ed.handlers = make(map[string][]EventHandlerInterface)
	
}