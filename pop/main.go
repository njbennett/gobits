package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/njbennett/gobits/pop/sims"
)

func main() {
	pop := sims.Population{
		&sims.Sim{ID: 0, Sex: 0, Born: 0},
		&sims.Sim{ID: 1, Sex: 1, Born: 0},
		&sims.Sim{ID: 2, Sex: 1, Born: 0},
	}

	gen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	limit, err := strconv.Atoi(os.Args[2])
	if err != nil {
		limit = -1
	}

	for i := 0; i < gen; i++ {
		pop = append(pop, pop.ThisYearsSims(i, limit)...)
	}

	for _, s := range pop {
		fmt.Println(s.Format())
	}
}
