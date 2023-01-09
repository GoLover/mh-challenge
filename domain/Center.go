package domain

const (
	AGENT_SCHEDULED EventType = iota
	AGENT_WALKED
	AGENT_RECEIVED
)

type EventType int
type Event struct {
	Type    EventType
	Message string
	AgentId int
}

type Center struct {
	Agents map[int]*Agent
}

type Coordinatation struct {
	XPOS int
	YPOS int
}

type CenterUsecase interface {
	SubmitCoordinate(Coordinatation)
	CoordinateLoop()
}
