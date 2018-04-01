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
	} else {
		return fmt.Sprintf("%d %d", s.parent1.id, s.parent2.id)
	}
}

func main() {
	sims := []sim{
		sim{id: 0, sex: "M"},
		sim{id: 1, sex: "F"},
	}

	sims = append(sims, sim{id: 2,
		sex:     "M",
		parent1: &sims[0],
		parent2: &sims[1],
	})

	sims = append(sims, sim{id: 3,
		sex:     "M",
		parent1: &sims[0],
		parent2: &sims[1],
	})

	fmt.Println(sims[0].format())
	fmt.Println(sims[1].format())
	fmt.Println(sims[2].format())
	fmt.Println(sims[3].format())
}
