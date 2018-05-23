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

type Population []*Sim

func NewSim(s0 *Sim, s1 *Sim, year int) (error, Sim) {
	if s0.Sex != 0 {
		return errors.New("parent0 should be sex 0"), Sim{}
	}

	if year-s0.Born >= 40 {
		msg := fmt.Sprintf("Sim ID %d is Age %d, too old to be parent0", s0.ID, s0.age(year))
		return errors.New(msg), Sim{}
	}

	if s0.Sex == s1.Sex {
		return errors.New("Parents cannot have the same sex"), Sim{}
	}

	if s1.death() < year {
		return errors.New("parent1 is too dead"), Sim{}
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
	return fmt.Sprintf("ID: %d Sex: %d Born: %d Died: %d Parents: %s", s.ID, s.Sex, s.Born, s.death(), parents)
}

func (s Sim) age(year int) int {
	return year - s.Born
}

func (s Sim) death() int {
	return s.Born + 80
}

func (s Sim) eligible(year int) bool {
	return s.Sex == 0 && s.age(year) >= 20 && s.age(year) <= 40
}

func (s Population) Eligible(year int) Population {
	eligible := Population{}
	for _, nextSim := range s {
		if nextSim.eligible(year) {
			eligible = append(eligible, nextSim)
		}
	}
	return eligible
}

func (s Population) ThisYearsSims(year int) Population {
	err, sim := NewSim(s[0], s[1], year)
	sim.ID = len(s)
	sim.Sex = len(s) % 2

	if err != nil {
		panic(err)
	}

	return Population{&sim}
}
