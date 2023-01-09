package delivery

import "mhlv/domain"

type Coordinatation struct {
	XPOS int `json:"XPOS"`
	YPOS int `json:"YPOS"`
}

func (c Coordinatation) ToDomain() domain.Coordinatation {
	return domain.Coordinatation{XPOS: c.XPOS, YPOS: c.YPOS}
}
