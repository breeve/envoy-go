package listenermanager

import (
	einit "github.com/breeve/envoy-go/envoy/init"
	"github.com/breeve/envoy-go/envoy/network"
	"github.com/breeve/envoy-go/envoy/server/common"
)

type ListenerImpl struct {
	dynamicInitManager *einit.ManagerImpl
	listenerManager    common.ListenerManagerHandler
}

func (l *ListenerImpl) AddSocketFactory(listenerFactory network.ListenSocketFactory) {

}

func (l *ListenerImpl) Initialize() {
	// 注册一个callback，通过 watcher 机制，挂到 worker 上，等待执行
	/* //
	.1
	local_init_watcher_(fmt::format("Listener-local-init-watcher {}", name),
						[this] {
						if (workers_started_) {
							parent_.onListenerWarmed(*this);
						} else {
							// Notify Server that this listener is
							// ready.
							listener_init_target_.ready();
						}
						}),
	.2
	if (workers_started_) {
		ENVOY_LOG_MISC(debug, "Initialize listener {} local-init-manager.", name_);
		// If workers_started_ is true, dynamic_init_manager_ should be initialized by listener
		// manager directly.
		dynamic_init_manager_->initialize(local_init_watcher_);
	}
	*/

}

func (l *ListenerImpl) RegisterToWorker() {
	l.listenerManager.OnListenerWarmed(l)
}

type ListenSocketFactoryImpl struct {
}
