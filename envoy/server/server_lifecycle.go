package server

type Stage string

const (
	Stage_Startup      = "Startup"
	Stage_PostInit     = "PostInit"
	Stage_ShutdownExit = "ShutdownExit"
)

type StageCallback func()

func ResgisterStageCallback(stage Stage, cb StageCallback) {

}
