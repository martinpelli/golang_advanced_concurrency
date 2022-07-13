package main

import "fmt"

type Editor struct {
	eventManager EventManager
}

func (Editor) openFile() {}

func (Editor) saveFile() {}

type EventManager struct {
	listeners []EventListeners
}

func (eventManager *EventManager) suscribe(listener EventListeners) {
	eventManager.listeners = append(eventManager.listeners, listener)
}

func (eventManager *EventManager) notify() {
	for _, listener := range eventManager.listeners {
		listener.update("myfile")
	}
}

type EventListeners interface {
	update(fileName string)
}

type EmailAlertListener struct{}

func (EmailAlertListener) update(fileName string) {
	fmt.Println("Do something with this file")
}

type LoggingListener struct{}

func (LoggingListener) update(fileName string) {
	fmt.Println("Do something with this file")
}

func main() {

	editor := Editor{
		eventManager: EventManager{},
	}

	email := EmailAlertListener{}
	logging := LoggingListener{}

	editor.eventManager.suscribe(email)
	editor.eventManager.suscribe(logging)

	editor.eventManager.notify()

}
