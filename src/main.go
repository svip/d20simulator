package main

import (
	"fmt"
	"strings"
)

type simStats struct {
	dc int
	runs int
	hits float64
}

func newStat(dc int) *simStats {
	return &simStats{dc, 0, 0}
}

func (ss *simStats) fraction() float64 {
	return (ss.hits/float64(ss.runs))
}

type sim struct {
	oldStats []*simStats
	newStats []*simStats
	mod int
	startDc int
	attempts int
}

func newSim(mod int, basedc int, dctries int) *sim {
	var oldStats []*simStats
	var newStats []*simStats
	for i := 0; i < dctries; i++ {
		oldStats = append(oldStats, newStat(basedc+i))
		newStats = append(newStats, newStat(basedc+i))
	}
	return &sim{
		oldStats,
		newStats,
		mod,
		basedc,
		dctries,
	}
}

func (s *sim) run(newRules bool, stats *simStats) {
	// TODO: Rewrite this to use the formula for the whole thing.
	for i := 1; i <= 20; i++ {
		if newRules && i <= s.mod {
			// The formula for whether a roll will succeed is as follows:
			// 20 - DC + mod + 1 / 20
			stats.hits += (20.0-float64(stats.dc)+float64(s.mod)+1.0)/20.0
		} else {
			if i == 20 {
				stats.hits++
			} else if i != 1 && i + s.mod >= stats.dc { // 1 always fails (not critical)
				stats.hits++
			}
		}
		stats.runs++
	}
}

func (s *sim) RunAll() {
	for _, simStat := range s.oldStats {
		s.run(false, simStat)
	}
	for _, simStat := range s.newStats {
		s.run(true, simStat)
	}
}

func (s *sim) PrintStats() {
	fmt.Printf("Ability mod: +%d\n", s.mod)
	fmt.Printf("DC\tStandard\tHouse\n")
	for i := 0; i < s.attempts; i++ {
		oldStat := s.oldStats[i]
		newStat := s.newStats[i]
		fmt.Printf("%d\t%.4f\t%.4f\n", oldStat.dc, oldStat.fraction(), newStat.fraction())
	}
	fmt.Println()
}

type simContainer struct {
	sims []*sim
}

func newSimContainer() simContainer{
	return simContainer{}
}

func (c *simContainer) addSim(mod int, basedc int, dctries int) {
	c.sims = append(c.sims, newSim(mod, basedc, dctries))
}

func (c *simContainer) RunAll() {
	for _, sim := range c.sims {
		sim.RunAll()
	}
}

func (c *simContainer) PrintAllStats() {
	for _, sim := range c.sims {
		sim.PrintStats()
	}
}

type point struct {
	x int
	y int
}

func (c *simContainer) DrawGraph() {
	width := 70
	pointWidth := 4
	height := 25
	for _, sim := range c.sims {
		fmt.Printf("+%d    * = standard, # = house, @ = both\n", sim.mod)
		var oldPoints []point
		var newPoints []point
		for _, stat := range sim.oldStats {
			x := (stat.dc-sim.startDc) * pointWidth
			y := height - int(stat.fraction()*float64(height))
			oldPoints = append(oldPoints, point{x,y})
		}
		for _, stat := range sim.newStats {
			x := (stat.dc-sim.startDc) * pointWidth
			y := height - int(stat.fraction()*float64(height))
			newPoints = append(newPoints, point{x,y})
		}
		fmt.Println()
		for h := 0; h < height; h++ {
			if h == 0 {
				fmt.Printf("100%% ")
			} else if h == height -1 {
				fmt.Printf("  0%% ")
			} else {
				fmt.Printf(strings.Repeat(" ", 5))
			}
			for w := 0; w < width; w+=pointWidth {
				foundPoint := []bool{false, false}
				for _, p := range oldPoints {
					if p.x == w && p.y == h {
						foundPoint[0] = true
					}
				}
				for _, p := range newPoints {
					if p.x == w && p.y == h {
						foundPoint[1] = true
					}
				}
				if foundPoint[0] && foundPoint[1] {
					fmt.Printf(" @")
				} else if foundPoint[0] {
					fmt.Printf(" *")
				} else if foundPoint[1] {
					fmt.Printf(" #")
				} else {
					fmt.Printf("  ")
				}
				fmt.Printf(strings.Repeat(" ", pointWidth-2))
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Printf(" DC: ")
		for _, stat := range sim.oldStats {
			dc := stat.dc
			if dc >= 10 {
				fmt.Printf("%d", dc)
			} else {
				fmt.Printf(" %d", dc)
			}
			fmt.Printf(strings.Repeat(" ", pointWidth-2))
		}
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
}

func main() {
	container := newSimContainer()

	container.addSim(0, 1, 18)
	container.addSim(1, 2, 18)
	container.addSim(2, 3, 18)
	container.addSim(3, 4, 18)
	container.addSim(4, 5, 18)
	container.addSim(5, 6, 18)
	container.addSim(6, 7, 18)
	container.addSim(7, 8, 18)
	container.addSim(8, 9, 18)
	container.addSim(9, 10, 18)
	container.addSim(10, 11, 18)

	container.RunAll()

	container.PrintAllStats()
	container.DrawGraph()
}
