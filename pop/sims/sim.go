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
	err := s0.canBeParent0(year)

	if err != nil {
		return err, Sim{}
	}

	err = s1.canBeParent1(year)

	if err != nil {
		return err, Sim{}
	}

	if s0.Parent1 == s1 {
		return errors.New("Parent1 cannot be the parent of Parent0"), Sim{}
	}

	return nil, Sim{
		Parent0: s0,
		Parent1: s1,
		Born:    year,
	}
}

func (s Sim) canBeParent0(year int) error {
	if s.Sex != 0 {
		msg := fmt.Sprintf("Sim ID %d is Sex %d, but should be Sex 0 to be parent0", s.ID, s.Sex)
		return errors.New(msg)
	}

	s0age := s.age(year)

	if s0age >= 40 {
		msg := fmt.Sprintf("Sim ID %d is Age %d, too old to be parent0", s.ID, s0age)
		return errors.New(msg)
	}

	if s0age <= 18 {
		msg := fmt.Sprintf("Sim ID %d is Age %d, too young to be parent0", s.ID, s0age)
		return errors.New(msg)
	}
	return nil
}

func (s Sim) canBeParent1(year int) error {
	if s.Sex != 1 {
		msg := fmt.Sprintf("Sim ID %d is Sex %d, but should be Sex 1 to be parent1", s.ID, s.Sex)
		return errors.New(msg)
	}

	s1age := s.age(year)

	if s1age <= 18 {
		msg := fmt.Sprintf("Sim ID %d is Age %d, too young to be parent0", s.ID, s1age)
		return errors.New(msg)
	}

	if s.death() < year {
		return errors.New("parent1 is too dead")
	}

	return nil
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

func (s Population) eligibleP1(year int) Population {
	newPop := Population{}

	for _, p := range s {
		if p.canBeParent1(year) == nil {
			newPop = append(newPop, p)
		}
	}
	return newPop
}

func (s Population) eligibleP0(year int) Population {
	newPop := Population{}

	for _, p := range s {
		if p.canBeParent0(year) == nil {
			newPop = append(newPop, p)
		}
	}
	return newPop
}

func (s Population) deadPopulation(year int) Population {
	newPop := Population{}

	for _, p := range s {
		if p.death() < year {
			newPop = append(newPop, p)
		}
	}
	return newPop
}

func (s Population) ThisYearsSims(year int, limit int) Population {
	genPop := Population{}
	popSize := len(s)
	deadPop := len(s.deadPopulation(year))

	parent1s := s.eligibleP1(year)
	parent0s := s.eligibleP0(year)

	for _, p0 := range parent0s {
		if limit > 0 && popSize-deadPop >= limit {
			break
		}
		for _, p1 := range parent1s {

			err, sim := NewSim(p0, p1, year)

			if err == nil {
				sim.ID = popSize
				sim.Sex = popSize % 2
				popSize++
				genPop = append(genPop, &sim)
				break
			}
		}

	}
	return genPop
}
