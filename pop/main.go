package main

import "fmt"

type sim struct {
	id      int
	sex     string
	parent1 *sim
	parent2 *sim
}

func (s sim) format() string {
	return fmt.Sprintf("ID: %d Sex: %s Parents: - -", s.id, s.sex)
}

func main() {
	sims := []sim{
		sim{id: 0, sex: "M"},
		sim{id: 1, sex: "F"},
	}

	fmt.Println(sims[0].format())
	fmt.Println(sims[1].format())
	fmt.Println("ID: 2 Sex: M Parents: 1 2")
	fmt.Println("ID: 3 Sex: F Parents: 1 2")
}
