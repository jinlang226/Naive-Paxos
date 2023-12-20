package paxos

import "testing"

func start(acceptorIds []int) []*Acceptor {
	//func start(acceptorIds []int, learnerIds []int) ([]*Acceptor, []*Learner) {

	acceptors := make([]*Acceptor, 0)
	for _, aid := range acceptorIds {
		a := newAcceptor(aid) //, learnerIds
		acceptors = append(acceptors, a)
	}

	//learners := make([]*Learner, 0)
	//for _, lid := range learnerIds {
	//	l := newLearner(lid, acceptorIds)
	//	learners = append(learners, l)
	//}

	return acceptors //, learners
}

func cleanup(acceptors []*Acceptor) { //, learners []*Learner
	for _, a := range acceptors {
		a.close()
	}

	//for _, l := range learners {
	//	l.close()
	//}
}

func TestSingleProposer(t *testing.T) {
	// 1001, 1002, 1003 是接受者 id
	acceptorIds := []int{1001, 1002, 1003}
	// 2001 是学习者 id
	//learnerIds := []int{2001}
	//acceptors, learns := start(acceptorIds, learnerIds)
	acceptors := start(acceptorIds)
	defer cleanup(acceptors) //, learns

	// 1 是提议者 id
	p := &Proposer{
		id:        1,
		acceptors: acceptorIds,
	}

	value := p.propose(6)
	if value != 6 {
		t.Errorf("value = %s, excepted %s", value, "hello world")
	}

	//learnValue := learns[0].chosen()
	//if learnValue != value {
	//	t.Errorf("learnValue = %s, excepted %s", learnValue, "hello world")
	//}
}
