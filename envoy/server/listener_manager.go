package server

import listenermanagerSource "github.com/breeve/envoy-go/envoy/source/extensions/listenermanagers/listenermanager"

//listenermanager "github.com/breeve/envoy/source/extensions/listener-managers/listener-manager"

type ListenerManager struct {
	workers []*Worker
}

func (m *ListenerManager) StartWorks() {

}

func (m *ListenerManager) AddOrUpdateListener() {
	m.addOrUpdateListenerInternal()
}

func (m *ListenerManager) addOrUpdateListenerInternal() {
	// .1 stop check

	// .2 warming 机制处理

	// .3 new
	// new_listener = std::make_unique<ListenerImpl>(config, version_info, *this, name, added_via_api, workers_started_, hash);
	newListener := &listenermanagerSource.ListenerImpl{}

	//
	m.setNewOrDrainingSocketFactory(newListener)

	//
	newListener.Initialize()
}

func (m *ListenerManager) setNewOrDrainingSocketFactory(listener *listenermanagerSource.ListenerImpl) {
	m.createListenSocketFactory(listener)
}

func (m *ListenerManager) createListenSocketFactory(listener *listenermanagerSource.ListenerImpl) {
	listener.AddSocketFactory(listenermanagerSource.ListenSocketFactoryImpl{})
}

func (m *ListenerManager) OnListenerWarmed(listener interface{}) {
	m.onListenerWarmed(listener.(*listenermanagerSource.ListenerImpl))
}

func (m *ListenerManager) onListenerWarmed(listener *listenermanagerSource.ListenerImpl) {
	for _, w := range m.workers {
		m.addListenerToWorker(w, listener)
	}
}

func (m *ListenerManager) addListenerToWorker(worker *Worker, listener *listenermanagerSource.ListenerImpl) {
	worker.AddListener(listener)
}
