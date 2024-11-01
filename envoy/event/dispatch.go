package event

type Dispatcher struct {
	stopCh            <-chan struct{}
	libeventScheduler *LibeventScheduler
}

func (d *Dispatcher) Run(rType RunType) {
	// .1 runPostCallbacks

	// .2
	d.libeventScheduler.Run(rType)

	<-d.stopCh
}

type RunType string

const (
	RunType_Block        RunType = "Block"
	RunType_NonBlock     RunType = "NonBlock"
	RunType_RunUntilExit RunType = "RunUntilExit"
)
