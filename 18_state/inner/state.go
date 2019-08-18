package inner

import "fmt"

type VoteState interface {
	Vote(user, item string, m VoteManager)
}

type NormalVoteState struct {
}

func (s *NormalVoteState) Vote(user, item string, m VoteManager) {
	m.MapVote[user] = item
	m.MapState[user] = &RepeatVoteState{}
	fmt.Println("恭喜投票成功")
}

type RepeatVoteState struct {
}

func (s *RepeatVoteState) Vote(user, item string, m VoteManager) {
	fmt.Println("请不要重复投票")
	if m.MapVoteCount[user] >= 4 {
		m.MapState[user] = &SpiteVoteState{}
	}
}

type SpiteVoteState struct {
}

func (s *SpiteVoteState) Vote(user, item string, m VoteManager) {
	if m.MapVote[user] != "" {
		delete(m.MapVote, user)
	}
	fmt.Println("你有刷票行为，取消投票资格")
	if m.MapVoteCount[user] >= 7 {
		m.MapState[user] = &BlackVoteState{}
	}
}

type BlackVoteState struct {
}

func (s *BlackVoteState) Vote(user, item string, m VoteManager) {
	fmt.Println("进入黑名单，禁止登陆使用本系统")
}
