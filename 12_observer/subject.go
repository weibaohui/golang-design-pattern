package observer

type Subject struct {
	observers []Observer
	content   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}
func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}
func (s *Subject) Detach(o Observer) {
	for i := 0; i < len(s.observers); i++ {
		if s.observers[i] == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			i--
		}
	}
}

func (s *Subject) notifyObservers() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Subject) SetContent(content string) {
	s.content = content
	s.notifyObservers()
}

func (s *Subject) GetContent() string {

	return s.content
}
