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
	WalkieTalkie               chan float64
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

func (a Agent) DistanceToCoordination(c Coordinatation) float64 {
	willingX := c.XPOS - a.Location.XPOS
	willingY := c.YPOS - a.Location.YPOS
	return math.Sqrt(math.Pow(float64(willingX), float64(2)) + math.Pow(float64(willingY), float64(2)))
}

func (a *Agent) Goto(c Coordinatation) {
	x := a.DistanceToCoordination(c)
	go func(remaindistance float64, reportingChan chan<- float64) {
		for {
			if remaindistance == float64(0) {
				break
			}
			time.Sleep(1 * time.Second)
			remaindistance -= float64(1)
			//fmt.Printf("goto remained,%f", remaindistance)
			reportingChan <- remaindistance
		}
	}(x, a.WalkieTalkie)
}
