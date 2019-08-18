package state

type VoteState interface {
	Vote(user, item string, m VoteManager)
}
