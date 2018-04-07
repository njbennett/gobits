package main

import (
	"fmt"

	"github.com/njbennett/gobits/pop/sims"
)

func eligible0(s []sims.Sim, year int) []*sims.Sim {
	var eligibleSims []*sims.Sim

	for _, nextSim := range s {
		if nextSim.Sex == 0 && year-nextSim.Born >= 20 {
			eligibleSims = append(eligibleSims, &nextSim)
		}
	}

	return eligibleSims
}

func main() {
	pop := []sims.Sim{
		sims.Sim{ID: 0, Sex: 0, Born: 0},
		sims.Sim{ID: 1, Sex: 1, Born: 0},
	}

	for i := 0; i < 30; i++ {
		parent1 := &pop[1]

		for _, parent0 := range eligible0(pop, i) {
			err, nextSim := sims.NewSim(parent0, parent1, i)
			if err != nil {
				break
			}
			nextSim.ID = len(pop)
			nextSim.Born = i
			nextSim.Sex = i % 2
			pop = append(pop, nextSim)
		}
	}

	for _, s := range pop {
		fmt.Println(s.Format())
	}
}
