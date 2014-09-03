package link

type Link interface {
	Run(cmd string) (string, error)

	Ready() bool
}
