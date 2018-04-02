package main

import "fmt"

type sim struct {
	id      int
	sex     int
	parent1 *sim
	parent2 *sim
	born    int
}

func (s sim) format() string {
	return fmt.Sprintf("ID: %d Born: %d Sex: %d Parents: %s", s.id, s.born, s.sex, s.parents())
}

func (s sim) parents() string {
	if s.parent1 == nil {
		return "- -"
	}

	return fmt.Sprintf("%d %d", s.parent1.id, s.parent2.id)
}

func newSim(s []sim, year int) (error, sim) {
	return nil, sim{
		id:      len(s),
		sex:     len(s) % 2,
		parent1: &s[0],
		parent2: &s[1],
		born:    year,
	}
}

func main() {
	sims := []sim{
		sim{id: 0, sex: 0, born: 0},
		sim{id: 1, sex: 1, born: 0},
	}

	for i := 20; i < 100; i++ {
		_, sim := newSim(sims, i)
		sims = append(sims, sim)
	}

	for _, s := range sims {
		fmt.Println(s.format())
	}
}
