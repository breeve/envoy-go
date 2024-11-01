package listenermanager

import "github.com/breeve/envoy-go/envoy/server/common"

type LdsApiImpl struct {
	listenerManager common.ListenerManagerHandler
}

func (l *LdsApiImpl) OnConfigUpdate() {
	//  for in adds, do one
	{ // one
		l.listenerManager.StartWorks()
	}
}
