package domain

type Center struct {
	Agents map[int]*Agent
}

type Coordinatation struct {
	XPOS int
	YPOS int
}

type CenterUsecase interface {
	CoordinateAgent(Coordinatation)
}
