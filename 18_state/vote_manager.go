package state

type VoteManager struct {
	voteState    VoteState
	MapVote      map[string]string
	mapVoteCount map[string]int
}

func NewVoteManager() *VoteManager {
	return &VoteManager{
		voteState:    nil,
		MapVote:      make(map[string]string, 0),
		mapVoteCount: make(map[string]int, 0),
	}
}

func (m *VoteManager) Vote(user, item string) {
	m.mapVoteCount[user] = m.mapVoteCount[user] + 1
	oldVoteCount := m.mapVoteCount[user]
	switch {
	case oldVoteCount == 1:
		m.voteState = &NormalVoteState{}
	case oldVoteCount > 1 && oldVoteCount < 5:
		m.voteState = &RepeatVoteState{}
	case oldVoteCount >= 5 && oldVoteCount < 8:
		m.voteState = &SpiteVoteState{}
	case oldVoteCount >= 8:
		m.voteState = &BlackVoteState{}
	}
	m.voteState.Vote(user, item, *m)
}
