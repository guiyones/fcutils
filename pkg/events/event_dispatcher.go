package events

import (
	"errors"
	"slices"
	"sync"
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

// Verifica se existe handler com o nome do evento passado
// Caso exista, vai percorrer todos os handlers e executa o metodo Handle deles com o evento dentro
// Vai pegar o Handler que foi registrado e executar o metodo Handle passando o evento que foi chamado
func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

// Vai verificar se existe registrado algum evento com o nome passado
// Caso exista ele vai percorrer os handlers de acordo com o nome do evento passado
// Para verificar se existe algum handler igual ao que estou tentando registrar
func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		if slices.Contains(ed.handlers[eventName], handler) {
			return ErrHandlerAlreadyRegistered
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

// O primeiro if verifica se o evento esta registrado
// Varifica se o handler que passou é o mesmo que foi registrado
func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		if slices.Contains(ed.handlers[eventName], handler) {
			return true
		}
	}
	return false
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for i := range ed.handlers[eventName] {
			if slices.Contains(ed.handlers[eventName], handler) {
				ed.handlers[eventName] = slices.Delete(ed.handlers[eventName], i, i+1)
				return nil
			}
		}
	}
	return nil
}

// Cria de novo uma struct zerada
func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}
