package upstream

import (
	"github.com/breeve/envoy_go/envoy/config"
)

type ClusterUpstream interface {
	AddOrUpdateCluster(cluster *config.Cluster, version string) bool
	SetPrimaryClustersInitializedCb(callback PrimaryClustersReadyCallback)
	SetInitializedCb(callback InitializationCompleteCallback)
	initializeSecondaryClusters()
}

type PrimaryClustersReadyCallback func()
type InitializationCompleteCallback func()
