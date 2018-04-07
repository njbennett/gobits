package sims

import (
	"errors"
	"fmt"
)

type Sim struct {
	ID      int
	Parent0 *Sim
	Parent1 *Sim
	Born    int
	Sex     int
}

func NewSim(s0 *Sim, s1 *Sim, year int) (error, Sim) {
	if year-s0.Born >= 40 {
		return errors.New("nope"), Sim{}
	}

	return nil, Sim{
		Parent0: s0,
		Parent1: s1,
		Born:    year,
	}
}

func (s Sim) Format() string {
	parents := "- -"
	if s.Parent0 != nil {
		parents = fmt.Sprintf("%d %d", s.Parent0.ID, s.Parent1.ID)
	}
	return fmt.Sprintf("ID: %d Sex: %d Born: %d Parents: %s", s.ID, s.Sex, s.Born, parents)
}
