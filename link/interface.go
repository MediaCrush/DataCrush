package link

type Link interface {
	func Connect(host string) error
	func Run(cmd string) (string, error)
	func Disconnect()
}
