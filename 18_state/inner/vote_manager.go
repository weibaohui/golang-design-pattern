package inner

type VoteManager struct {
	MapVote      map[string]string
	MapVoteCount map[string]int
	MapState     map[string]VoteState
}

func NewVoteManager() *VoteManager {
	return &VoteManager{
		MapVote:      make(map[string]string, 0),
		MapVoteCount: make(map[string]int, 0),
		MapState:     make(map[string]VoteState, 0),
	}
}

func (m *VoteManager) Vote(user, item string) {
	m.MapVoteCount[user] = m.MapVoteCount[user] + 1
	state := m.MapState[user]
	if state == nil {
		state = &NormalVoteState{}
	}
	state.Vote(user, item, *m)
}
