package server

import (
	"github.com/breeve/envoy-go/envoy/event"
	listenermanagerSource "github.com/breeve/envoy-go/envoy/source/extensions/listenermanagers/listenermanager"
)

type Worker struct {
	dispatcher *event.Dispatcher
	connection *listenermanagerSource.ConnectionHandlerImpl
}

func (w *Worker) Start() {
	go w.threadRoutine()
}

func (w *Worker) threadRoutine() {
	w.dispatcher.Run(event.RunType_Block)
}

func (w *Worker) AddListener(listener *listenermanagerSource.ListenerImpl) {
	w.connection.AddListener(listener)
}

type WorkerFactory struct {
}

func (f *WorkerFactory) CreateWorker() *Worker {
	return nil
}
