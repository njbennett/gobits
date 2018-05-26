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

	simsch := make(chan *sims.Sim)
	go printSims(simsch)

	for i := 0; i < gen; i++ {
		thisYearsSims := pop.ThisYearsSims(i, limit)
		pop = append(pop, thisYearsSims...)
		feed(thisYearsSims, simsch)
	}
}

func printSims(ch chan *sims.Sim) {
	for {
		s := <-ch
		fmt.Println(s.Format())
	}
}

func feed(p []*sims.Sim, ch chan *sims.Sim) {
	for _, s := range p {
		ch <- s
	}
}
