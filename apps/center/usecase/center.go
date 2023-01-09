package usecase

import (
	"fmt"
	"mhlv/domain"
)

type CenterUsecase struct {
	center             *domain.Center
	jobQueue           chan domain.Coordinatation
	agentEvents        chan domain.Event
	continueJobProcess chan bool
	availableAgents    int
}

func NewCenterUsecase(nubmerOfAgents int) *CenterUsecase {
	cu := &CenterUsecase{continueJobProcess: make(chan bool), jobQueue: make(chan domain.Coordinatation, 1), agentEvents: make(chan domain.Event), availableAgents: nubmerOfAgents}
	agents := make(map[int]*domain.Agent, nubmerOfAgents)
	i := 0
	for i < nubmerOfAgents {
		agents[i] = &domain.Agent{Id: i,
			Location:  domain.Coordinatation{XPOS: 0, YPOS: 0},
			Available: true, NextLocationRemainingSteps: 0, WalkieTalkie: cu.agentEvents}
		i++
	}
	cu.center = &domain.Center{Agents: agents}
	return cu
}

func (cu *CenterUsecase) SubmitCoordinate(nc domain.Coordinatation) {
	cu.jobQueue <- nc
}
func (cu *CenterUsecase) CoordinateLoop() {
	go cu.agentloop()
	go cu.jobdispatchLoop()
}
func (cu *CenterUsecase) agentloop() {
	for {
		agentEvent := <-cu.agentEvents
		switch agentEvent.Type {
		case domain.AGENT_RECEIVED:
			cu.availableAgents += 1
			err := cu.center.Agents[agentEvent.AgentId].MakeItAvailable()
			if err != nil {
				panic(err)
			}
			fmt.Println(agentEvent.Message)
			go func() { cu.continueJobProcess <- true }()
		case domain.AGENT_SCHEDULED:
			cu.availableAgents -= 1
			cu.center.Agents[agentEvent.AgentId].MakeItBusy()
			fmt.Println(agentEvent.Message)
		case domain.AGENT_WALKED:
			//fmt.Println(agentEvent.Message)
		}
	}
}

func (cu *CenterUsecase) jobdispatchLoop() {
	for {
		jobDispatch := <-cu.jobQueue
		if cu.availableAgents == 0 {
			<-cu.continueJobProcess
		}
		availableAgents := make([]*domain.Agent, 0)
		for len(availableAgents) == 0 {
			for _, k := range cu.center.Agents {
				if k.IsAgentAvailable() {
					k.MakeItBusy()
					availableAgents = append(availableAgents, k)
				}
			}
		}

		var nearestAgentId = -1
		var distaneOfAgent = 0
		for _, k := range availableAgents {
			dtc := k.DistanceToCoordination(jobDispatch)
			if dtc == 0 {
				nearestAgentId = -1
				fmt.Printf("agent number: %d is already in location.\n", k.Id)
				break
			}
			if nearestAgentId == -1 {
				nearestAgentId = k.Id
			}
			if dtc < distaneOfAgent {
				distaneOfAgent = dtc
				nearestAgentId = k.Id
			}
		}

		for _, k := range availableAgents {
			if k.Id != nearestAgentId {
				k.MakeItAvailable()
			}
		}
		if nearestAgentId == -1 {
			continue
		}
		cu.center.Agents[nearestAgentId].Goto(jobDispatch)
	}
}
