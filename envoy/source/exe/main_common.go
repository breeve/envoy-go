package exe

import "github.com/breeve/envoy-go/envoy/server"

type StrippedMainBase struct {
	option Option
	server *server.InstanceImpl
}
type Option string

const (
	Option_Serve    Option = "Serve"
	Option_Validate Option = "Validate"
	Option_InitOnly Option = "InitOnly"
)

type MainCommonBase struct {
	stripped *StrippedMainBase
}

type MainCommon struct {
	base *MainCommonBase
}

func (m *MainCommon) Server() {
}

func (m *MainCommon) Run() {
	switch m.base.stripped.option {
	case Option_Serve:
		m.runServer()
	case Option_Validate:
	case Option_InitOnly:
	}
}

func (m *MainCommon) runServer() {
	m.base.stripped.server.Run()
}

func MainThread() {
	mainCommon := MainCommon{}
	mainCommon.Run()
}
