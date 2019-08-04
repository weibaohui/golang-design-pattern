package pc

type Mediator interface {
	Change(colleague Colleague)
}
