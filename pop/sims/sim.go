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
		msg := fmt.Sprintf("Sim ID %d is Sex %d, but should be Sex 0 to be parent0", s0.ID, s0.Sex)
		return errors.New(msg), Sim{}
	}

	s0age := s0.age(year)
	if s0age >= 40 {
		msg := fmt.Sprintf("Sim ID %d is Age %d, too old to be parent0", s0.ID, s0age)
		return errors.New(msg), Sim{}
	}

	if s0age <= 18 {
		msg := fmt.Sprintf("Sim ID %d is Age %d, too young to be parent0", s0.ID, s0age)
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
	p1 := s[1]
	pop := Population{}

	for _, p0 := range s {
		err, sim := NewSim(p0, p1, year)

		if err == nil {
			sim.ID = (len(s))
			sim.Sex = (len(s)) % 2
			pop = append(pop, &sim)
		}
	}
	return pop
}
