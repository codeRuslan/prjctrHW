package observer

type Subject interface {
	Deregister(observer Observer)
	Register(observer Observer)
	NotifyAll()
}
