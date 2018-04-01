package main

import "fmt"

type sim struct {
	id      int
	sex     string
	parent1 *sim
	parent2 *sim
}

func (s sim) format() string {
	return fmt.Sprintf("ID: %d Sex: %s Parents: %s", s.id, s.sex, s.parents())
}

func (s sim) parents() string {
	if s.parent1 == nil {
		return "- -"
	}

	return fmt.Sprintf("%d %d", s.parent1.id, s.parent2.id)
}

func newSim(s []sim) sim {
	sex := "F"
	if len(s)%2 == 0 {
		sex = "M"
	}
	return sim{
		id:      len(s),
		sex:     sex,
		parent1: &s[0],
		parent2: &s[1],
	}
}

func main() {
	sims := []sim{
		sim{id: 0, sex: "M"},
		sim{id: 1, sex: "F"},
	}

	sims = append(sims, newSim(sims))
	sims = append(sims, newSim(sims))

	for _, s := range sims {
		fmt.Println(s.format())
	}
}
