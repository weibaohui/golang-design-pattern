package builder

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct(header, body, footer string) {
	d.builder.BuildHeader(header)
	d.builder.BuildBody(body)
	d.builder.BuildFooter(footer)
}
