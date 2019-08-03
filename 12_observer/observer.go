package observer

type Observer interface {
	Update(subject *Subject)
}
