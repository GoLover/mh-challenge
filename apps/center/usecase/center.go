package usecase

import (
	"fmt"
	"mhlv/domain"
)

type CenterUsecase struct {
	center *domain.Center
}

func NewCenterUsecase(nubmerOfAgents int) *CenterUsecase {
	agents := make(map[int]*domain.Agent, nubmerOfAgents)
	i := 0
	for i < nubmerOfAgents {
		agents[i] = &domain.Agent{Id: i,
			Location:  domain.Coordinatation{XPOS: 0, YPOS: 0},
			Available: true, NextLocationRemainingSteps: 0, WalkieTalkie: make(chan float64)}
		i++
	}
	return &CenterUsecase{center: &domain.Center{Agents: agents}}
}

func (cu *CenterUsecase) CoordinateAgent(nextCoordination domain.Coordinatation) {
	availableAgents := make([]*domain.Agent, 0)
	for len(availableAgents) == 0 {
		//fmt.Println("first for")
		//fmt.Println(len(cu.center.Agents))
		for _, k := range cu.center.Agents {
			//fmt.Println("second for")
			if k.IsAgentAvailable() {
				k.MakeItBusy()
				availableAgents = append(availableAgents, k)
			}
		}
	}

	var nearestAgentId = -1
	var distaneOfAgent float64 = float64(0)
	//fmt.Println("length of available agents: ")
	//fmt.Println(len(availableAgents))
	for _, k := range availableAgents {
		//fmt.Println("third for")
		dtc := k.DistanceToCoordination(nextCoordination)
		//fmt.Printf("number of dtc %f", dtc)
		if nearestAgentId == -1 {
			distaneOfAgent = dtc
			nearestAgentId = k.Id
		}
		if dtc < distaneOfAgent {
			distaneOfAgent = dtc
			nearestAgentId = k.Id
		}
	}
	//fmt.Printf("dta determined as: %f, and choosen agent is: %d", distaneOfAgent, nearestAgentId)
	for _, k := range availableAgents {
		//fmt.Println("fifth for")
		if k.Id != nearestAgentId {
			k.MakeItAvailable()
		}
	}
	//fmt.Println(nearestAgentId)
	go cu.center.Agents[nearestAgentId].Goto(nextCoordination)
	for {
		//fmt.Println("sixth for")
		remaining := <-cu.center.Agents[nearestAgentId].WalkieTalkie
		fmt.Printf("remaining steps: %f, for agent: %d \n", remaining, nearestAgentId)
		if remaining == 0 {
			break
		}
	}
	err := cu.center.Agents[nearestAgentId].MakeItAvailable()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Agent %d arrived to coordination %#v \n", nearestAgentId, nextCoordination)
}
