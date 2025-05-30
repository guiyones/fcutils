package events

import (
	"sync"
	"time"
)

// Definição do evento
type EventInterface interface {
	GetName() string        // Para recuperar os dados com o nome do evento
	GetDateTime() time.Time // Data e hora que o evento foi disparado
	GetPayLoad() any        // Dados que tem no evento. Pode ter varios payloads em varios formatos
}

// Operações que serão executadas quando um evento é chamado
// Chamou o evento o Handler executa
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup) // É o metodo que executa a operação, para isso precisa do evento
}

// Gerenciado de eventos
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error // Registra um evento
	Dispatch(event EventInterface) error                            // Faz com que os eventos aconteçam e que os handlres sejam executados
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool // Verifica se tem um event name com o handler
	Clear() error                                             // Limpa o event dispatcher e mata todos os eventos que estão registrados
}
