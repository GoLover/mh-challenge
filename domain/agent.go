package domain

import (
	"fmt"
	"math"
)

type Agent struct {
	Id                         int
	Location                   Coordinatation
	Available                  bool
	NextLocationRemainingSteps int
	WalkieTalkie               chan Event
}

func (a *Agent) MakeItBusy() {
	if a.IsAgentAvailable() {
		a.Available = false
	}
}

func (a *Agent) MakeItAvailable() error {
	if a.NextLocationRemainingSteps == 0 {
		a.Available = true
		return nil
	}
	return ErrAgentIsOnTheWay
}
func (a Agent) IsAgentAvailable() bool {
	return a.Available
}

func (a Agent) DistanceToCoordination(c Coordinatation) int {
	willingX := c.XPOS - a.Location.XPOS
	willingY := c.YPOS - a.Location.YPOS
	return int(math.Sqrt(math.Pow(float64(willingX), float64(2)) + math.Pow(float64(willingY), float64(2))))
}

func (a *Agent) Goto(c Coordinatation) {
	x := a.DistanceToCoordination(c)
	a.WalkieTalkie <- Event{Type: AGENT_SCHEDULED, Message: fmt.Sprintf("agent %d scheduled for coordination of %#v", a.Id, c), AgentId: a.Id}
	go func(remaindistance int, reportingChan chan<- Event) {
		for {
			reportingChan <- Event{Type: AGENT_WALKED, Message: fmt.Sprintf("remained distance %d, for agent %d", remaindistance, a.Id), AgentId: a.Id}
			if remaindistance == int(0) {
				a.Location = c
				reportingChan <- Event{Type: AGENT_RECEIVED, AgentId: a.Id, Message: fmt.Sprintf("Agent %d arrived to coordination %#v", a.Id, c)}
				break
			}
			//time.Sleep(1 * time.Second)
			remaindistance -= int(1)
		}
	}(x, a.WalkieTalkie)
}
