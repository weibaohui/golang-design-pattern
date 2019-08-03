package observer

import "fmt"

type reader struct {
	Name string
}

func (r *reader) Update(s *Subject) {
	fmt.Printf("%s收到报纸，内容是==%s \n", r.Name, s.GetContent())
}
