package domain

import (
	"math"
	"time"
)

type Agent struct {
	Id                         int
	Location                   Coordinatation
	Available                  bool
	NextLocationRemainingSteps int
	WalkieTalkie               chan int
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
	go func(remaindistance int, reportingChan chan<- int) {
		for {
			reportingChan <- remaindistance
			if remaindistance == int(0) {
				a.Location = c
				break
			}
			time.Sleep(1 * time.Second)
			remaindistance -= int(1)
			//fmt.Printf("goto remained,%d", remaindistance)

		}
	}(x, a.WalkieTalkie)
}
