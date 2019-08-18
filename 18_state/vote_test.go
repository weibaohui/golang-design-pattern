package state

import (
	"testing"
)

func TestRun(t *testing.T) {
	voteManager := NewVoteManager()
	for i := 0; i < 8; i++ {
		voteManager.Vote("user1", "a,b,c")
	}
}
