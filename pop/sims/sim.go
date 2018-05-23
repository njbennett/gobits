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
	Died    int
}

type Population []*Sim

func NewSim(s0 *Sim, s1 *Sim, year int) (error, Sim) {
	if year-s0.Born >= 40 {
		return errors.New("nope"), Sim{}
	}

	if s0.Sex == s1.Sex {
		return errors.New("Parents cannot have the same sex"), Sim{}
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

func (s Population) Eligible(year int) Population {
	eligible := Population{}
	for _, nextSim := range s {
		if nextSim.Sex == 0 && year-nextSim.Born >= 20 && year-nextSim.Born <= 40 {
			eligible = append(eligible, nextSim)
		}
	}
	return eligible
}

func (s Population) Cull(year int) Population {
	for _, nextSim := range s {
		if year-nextSim.Born >= 80 {
			nextSim.Died = year
		}
	}
	return s
}
