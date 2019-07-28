package builder

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	var builder Builder
	builder = &TxtBuilder{}
	Run(builder)
	builder = &XmlBuilder{}
	Run(builder)
}
func Run(builder Builder) {
	head, body, footer := "head", "body", "footer"
	director := NewDirector(builder)
	director.Construct(head, body, footer)
	fmt.Println(builder.Result())
}
