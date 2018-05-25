package sims_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/njbennett/gobits/pop/sims"
)

var _ = Describe("Sims", func() {

	Describe("NewSim", func() {
		It("makes a new sim", func() {
			s0 := Sim{Born: 0, Sex: 0}
			s1 := Sim{Born: 0, Sex: 1}
			year := 20

			expectedSim := Sim{
				Parent0: &s0,
				Parent1: &s1,
				Born:    year,
			}

			err, newSim := NewSim(&s0, &s1, year)

			Expect(err).NotTo(HaveOccurred())
			Expect(newSim).To(Equal(expectedSim))
		})

		It("returns an error when parent0 is too old", func() {
			s0 := Sim{ID: 0, Sex: 0, Born: 0}
			s1 := Sim{ID: 1, Sex: 1, Born: 0}
			year := 40

			err, newSim := NewSim(&s0, &s1, year)
			Expect(err).To(Equal(errors.New("Sim ID 0 is Age 40, too old to be parent0")))
			Expect(newSim).To(Equal(Sim{}))
		})

		It("returns an error when parent1 is parent of parent0", func() {
			s0 := Sim{Sex: 0, Born: 0}
			s1 := Sim{Sex: 1, Born: 0}
			s2 := Sim{Sex: 0, Born: 20, Parent1: &s1, Parent0: &s0}
			year := 40

			err, newSim := NewSim(&s2, &s1, year)
			Expect(err).To(Equal(errors.New("Parent1 cannot be the parent of Parent0")))
			Expect(newSim).To(Equal(Sim{}))
		})

		It("returns an error when either parent is dead", func() {
			s0 := Sim{Sex: 0, Born: 60}
			s1 := Sim{Sex: 1, Born: 0}
			year := 90

			err, newSim := NewSim(&s0, &s1, year)
			Expect(err).To(Equal(errors.New("parent1 is too dead")))
			Expect(newSim).To(Equal(Sim{}))
		})

		It("returns an error when parent0 is not sex 0", func() {
			s0 := Sim{ID: 0, Born: 0, Sex: 0}
			s1 := Sim{ID: 1, Born: 0, Sex: 1}
			year := 20

			err, newSim := NewSim(&s1, &s0, year)
			Expect(err).To(Equal(errors.New("Sim ID 1 is Sex 1, but should be Sex 0 to be parent0")))
			Expect(newSim).To(Equal(Sim{}))
		})

		It("returns an error when parent0 is too young", func() {
			s0 := Sim{ID: 0, Born: 10, Sex: 0}
			s1 := Sim{ID: 1, Born: 0, Sex: 1}
			year := 20

			err, newSim := NewSim(&s0, &s1, year)
			Expect(err).To(Equal(errors.New("Sim ID 0 is Age 10, too young to be parent0")))
			Expect(newSim).To(Equal(Sim{}))
		})
	})

	Describe("Format", func() {
		It("formats Sim for printing", func() {
			noParents0 := "ID: 0 Sex: 0 Born: 0 Died: 80 Parents: - -"
			noParents1 := "ID: 1 Sex: 1 Born: 0 Died: 80 Parents: - -"
			hasParents := "ID: 2 Sex: 0 Born: 20 Died: 100 Parents: 0 1"

			sim0 := Sim{Born: 0, ID: 0, Sex: 0}
			sim1 := Sim{Born: 0, ID: 1, Sex: 1}
			sim2 := Sim{Born: 20, ID: 2, Sex: 0, Parent0: &sim0, Parent1: &sim1}

			Expect(sim0.Format()).To(Equal(noParents0))
			Expect(sim1.Format()).To(Equal(noParents1))
			Expect(sim2.Format()).To(Equal(hasParents))
		})
	})

	Describe("ThisYearsSims", func() {
		It("generates a new batch of sims", func() {
			year := 21
			population := Population{
				&Sim{ID: 0, Sex: 0, Born: 0},
				&Sim{ID: 1, Sex: 1, Born: 0},
			}

			_, sim := NewSim(population[0], population[1], year)
			sim.ID = len(population)

			thisYearsSims := Population{&sim}

			Expect(population.ThisYearsSims(year)).To(Equal(thisYearsSims))
		})

		It("assigns sex based on population size", func() {
			year := 22
			population := Population{
				&Sim{ID: 0, Sex: 0, Born: 0},
				&Sim{ID: 1, Sex: 1, Born: 0},
				&Sim{ID: 2, Sex: 0, Born: 0},
			}

			firstBorn := population.ThisYearsSims(year)[0]

			Expect(firstBorn.Sex).To(Equal(1))
		})

		It("generates one child for every eligible Sex 0 sim", func() {
			year := 50
			population := Population{
				&Sim{ID: 0, Sex: 0, Born: 0},
				&Sim{ID: 1, Sex: 1, Born: 0},
				&Sim{ID: 2, Sex: 0, Born: 0},
				&Sim{ID: 3, Sex: 0, Born: 20},
				&Sim{ID: 4, Sex: 0, Born: 20},
				&Sim{ID: 5, Sex: 0, Born: 40},
			}

			Expect(len(population.ThisYearsSims(year))).To(Equal(2))
		})

		It("includes generated sims in population size for sex and ID", func() {
			year := 50
			population := Population{
				&Sim{ID: 0, Sex: 0, Born: 0},
				&Sim{ID: 1, Sex: 1, Born: 0},
				&Sim{ID: 2, Sex: 0, Born: 0},
				&Sim{ID: 3, Sex: 0, Born: 20},
				&Sim{ID: 4, Sex: 0, Born: 20},
				&Sim{ID: 5, Sex: 0, Born: 40},
			}

			Expect(population.ThisYearsSims(year)[0].ID).To(Equal(6))
			Expect(population.ThisYearsSims(year)[0].Sex).To(Equal(0))
			Expect(population.ThisYearsSims(year)[1].ID).To(Equal(7))
			Expect(population.ThisYearsSims(year)[1].Sex).To(Equal(1))
		})

		It("selects an eligible p1 for each p0, if possible", func() {
			year := 30
			population := Population{
				&Sim{ID: 0, Sex: 1, Born: 0},
				&Sim{ID: 1, Sex: 1, Born: 0},
			}

			population = append(population, &Sim{ID: 3, Sex: 0, Born: 0, Parent1: population[0]})

			Expect(len(population.ThisYearsSims(year))).To(Equal(1))
			parent1 := population.ThisYearsSims(year)[0].Parent1
			expectedParent1 := population[1]
			Expect(parent1).To(Equal(expectedParent1))
		})

		Measure("it should handle populations with many dead sims quickly", func(b Benchmarker) {
			runtime := b.Time("runtime", func() {
				year := 100
				population := Population{
					&Sim{ID: 0, Sex: 0, Born: 80},
					&Sim{ID: 1, Sex: 1, Born: 80},
				}

				for i := 2; i < 1000; i++ {
					population = append(population, &Sim{ID: i, Sex: 0, Born: 0})
				}
				for i := 1001; i < 100000; i++ {
					population = append(population, &Sim{ID: i, Sex: 1, Born: 0})
				}

				Expect(len(population.ThisYearsSims(year))).To(Equal(1))
			})
			Expect(runtime.Seconds()).Should(BeNumerically("<", 0.2))
		}, 10)

	})
})
