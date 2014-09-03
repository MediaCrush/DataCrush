package agent

type Event struct {
	Agent   string
	Source  string
	Payload interface{}
}

type CPUEvent struct {
	Value     float64
	Timestamp int64
}
