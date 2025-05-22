package events

import (
	"errors"
	"slices"
)

var ErrHandlerAlreadyRegistered error = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {

	// if _, ok := ed.handlers[eventName]; ok {
	// 	for _, h := range ed.handlers[eventName] {
	// 		if h == handler {
	// 			return ErrHandlerAlreadyRegistered
	// 		}
	// 	}
	// }

	// Vai verificar se existe registrado algum evento com o nome passado
	// Caso exista ele vai percorrer os handlers de acordo com o nome do evento passado
	// Para verificar se existe algum handler igual ao que estou tentando registrar
	if _, ok := ed.handlers[eventName]; ok {
		if slices.Contains(ed.handlers[eventName], handler) {
			return ErrHandlerAlreadyRegistered
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}
