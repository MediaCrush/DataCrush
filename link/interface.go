package link

type Link interface {
	Connect(host string) error
	Run(cmd string) (string, error)
	Disconnect()

	Ready() bool
}
