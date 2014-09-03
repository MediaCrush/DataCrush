package agent

type Agent interface {
	Run(events chan<- Event)
	Stop()
}
