package faketime

type Observer interface {
	Update() error
}
