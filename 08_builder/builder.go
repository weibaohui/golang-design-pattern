package builder

import "fmt"

type Builder interface {
	BuildHeader(header string)
	BuildBody(body string)
	BuildFooter(footer string)
	Result() string
}

type TxtBuilder struct{ result string }

func (b *TxtBuilder) BuildHeader(str string) {
	b.result = fmt.Sprintf("%s TxtBuilder header=%s; \n", b.result, str)
}

func (b *TxtBuilder) BuildBody(str string) {
	b.result = fmt.Sprintf("%s TxtBuilder body=%s; \n", b.result, str)
}

func (b *TxtBuilder) BuildFooter(str string) {
	b.result = fmt.Sprintf("%s TxtBuilder footer=%s;\n", b.result, str)
}

func (b *TxtBuilder) Result() string {
	return b.result
}

type XmlBuilder struct{ result string }

func (b *XmlBuilder) BuildHeader(str string) {
	b.result = fmt.Sprintf("%s XmlBuilder header=%s; \n", b.result, str)
}

func (b *XmlBuilder) BuildBody(str string) {
	b.result = fmt.Sprintf("%s XmlBuilder body=%s; \n", b.result, str)
}

func (b *XmlBuilder) BuildFooter(str string) {
	b.result = fmt.Sprintf("%s XmlBuilder footer=%s;\n", b.result, str)
}

func (b *XmlBuilder) Result() string {
	return b.result
}
