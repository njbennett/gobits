package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/njbennett/gobits/pop/sims"
)

func main() {
	pop := sims.Population{
		&sims.Sim{ID: 0, Sex: 0, Born: 0, Died: 80},
		&sims.Sim{ID: 1, Sex: 1, Born: 0, Died: 80},
	}

	gen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	for i := 0; i < gen; i++ {
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
