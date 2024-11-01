package listenermanager

type ConnectionHandlerImpl struct {
	listeners []ActiveListener
}

func (c *ConnectionHandlerImpl) AddListener(listener *ListenerImpl) {
	// UDP

	// TCP
	tcp := ActiveTcpListener{}
	c.listeners = append(c.listeners, tcp)

}
