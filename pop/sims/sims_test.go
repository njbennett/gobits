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
			s0 := Sim{Sex: 0, Born: 0}
			s1 := Sim{Sex: 1, Born: 0}
			year := 40

			err, newSim := NewSim(&s0, &s1, year)
			Expect(err).To(Equal(errors.New("parent0 is too old")))
			Expect(newSim).To(Equal(Sim{}))
		})

		It("returns an error when parents are the same sex", func() {
			s0 := Sim{Sex: 0, Born: 0}
			s1 := Sim{Sex: 0, Born: 0}
			year := 20

			err, newSim := NewSim(&s0, &s1, year)
			Expect(err).To(Equal(errors.New("Parents cannot have the same sex")))
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

	Describe("Eligible", func() {
		It("returns all sex 0 sims over 20 and under 40", func() {
			year := 50
			population := Population{
				&Sim{Sex: 1, Born: 0},
				&Sim{Sex: 0, Born: 0},
				&Sim{Sex: 0, Born: 20},
				&Sim{Sex: 0, Born: 40},
			}

			eligiblePopulation := Population{
				&Sim{Sex: 0, Born: 20},
			}

			Expect(population.Eligible(year)).To(Equal(eligiblePopulation))
		})
	})

	Describe("ThisYearsSims", func() {
		It("generates a new batch of sims", func() {
			year := 21
			population := Population{
				&Sim{ID: 0, Sex: 1, Born: 0},
				&Sim{ID: 1, Sex: 0, Born: 0},
			}

			_, sim := NewSim(population[0], population[1], year)
			sim.ID = len(population)

			thisYearsSims := Population{&sim}

			Expect(population.ThisYearsSims(year)).To(Equal(thisYearsSims))
		})

		It("assigns sex based on population size", func() {
			year := 22
			population := Population{
				&Sim{ID: 0, Sex: 1, Born: 0},
				&Sim{ID: 1, Sex: 0, Born: 0},
				&Sim{ID: 2, Sex: 0, Born: 0},
			}

			_, sim := NewSim(population[0], population[1], year)
			sim.ID = len(population)
			sim.Sex = len(population) % 2

			thisYearsSims := Population{&sim}

			Expect(population.ThisYearsSims(year)).To(Equal(thisYearsSims))
		})

	})
})
