package observer

type Observer interface {
	Update(string, string)
	GetID() string
}
