package faketime

type Subject interface {
	Register(observer Observer)
	//deregister(observer Observer)
	NotifyAll()
}
