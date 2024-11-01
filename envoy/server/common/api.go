package common

type ListenerManagerHandler interface {
	StartWorks()
	OnListenerWarmed(listener interface{})
}
