package obs

type Observer interface {
	Update(string)
	GetID() string
}

type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}
