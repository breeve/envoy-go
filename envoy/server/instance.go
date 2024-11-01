package server

import (
	"github.com/breeve/envoy-go/envoy/accesslog"
	"github.com/breeve/envoy-go/envoy/event"
	einit "github.com/breeve/envoy-go/envoy/init"
	"github.com/breeve/envoy-go/envoy/network"
	"github.com/breeve/envoy-go/envoy/secret"
	"github.com/breeve/envoy-go/envoy/server/overload"
	"github.com/breeve/envoy-go/envoy/ssl"
	"github.com/breeve/envoy-go/envoy/upstream"
)

type Instance struct {
	clusterManager    *upstream.ClusterManager
	sslContextManager *ssl.ContextManager
	dispatcher        *event.Dispatcher
	dnsResolver       network.DnsResolver
	drainManager      *DrainManager
	accessLogManager  *accesslog.AccessLogManager
	hotRestart        *HotRestart
	initManager       *einit.Manager
	listenerManager   *ListenerManager
	overloadManager   *overload.OverloadManager
	secretManager     *secret.SecretManager
}

type InstanceImpl struct {
	Instance *Instance
}

func (i *InstanceImpl) Run() {
	// .1
	ResgisterStageCallback(Stage_PostInit, i.startWorkers)

	// .2
	// auto watchdog = main_thread_guard_dog_->createWatchDog(api_->threadFactory().currentThreadId(),
	// "main_thread", *dispatcher_);

	// .3
	// dispatcher_->post([this] { notifyCallbacksForStage(Stage::Startup); });

	// .4
	i.Instance.dispatcher.Run(event.RunType_Block)

}

func (i *InstanceImpl) startWorkers() {
	i.Instance.listenerManager.StartWorks()
}
