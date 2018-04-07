package main

import (
	"fmt"

	"github.com/njbennett/gobits/pop/sims"
)

func main() {
	pop := sims.Population{
		&sims.Sim{ID: 0, Sex: 0, Born: 0},
		&sims.Sim{ID: 1, Sex: 1, Born: 0},
	}

	for i := 0; i < 21; i++ {
		parent1 := pop[1]

		for _, parent0 := range pop.Eligible(i) {
			err, nextSim := sims.NewSim(parent0, parent1, i)
			if err != nil {
				break
			}
			nextSim.ID = len(pop)
			nextSim.Born = i
			nextSim.Sex = i % 2
			pop = append(pop, &nextSim)
		}
	}

	for _, s := range pop {
		fmt.Println(s.Format())
	}
}
