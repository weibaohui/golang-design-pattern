package state

import "fmt"

type NormalVoteState struct {
}

func (s *NormalVoteState) Vote(user, item string, m VoteManager) {
	m.MapVote[user] = item
	fmt.Println("恭喜投票成功")
}

type RepeatVoteState struct {
}

func (s *RepeatVoteState) Vote(user, item string, m VoteManager) {
	fmt.Println("请不要重复投票")
}

type SpiteVoteState struct {
}

func (s *SpiteVoteState) Vote(user, item string, m VoteManager) {
	if m.MapVote[user] != "" {
		delete(m.MapVote, user)
	}
	fmt.Println("你有刷票行为，取消投票资格")
}

type BlackVoteState struct {
}

func (s *BlackVoteState) Vote(user, item string, m VoteManager) {
	fmt.Println("进入黑名单，禁止登陆使用本系统")
}
