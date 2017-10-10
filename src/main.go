package main

import (
	"math/rand"
	"time"
	"fmt"
	"strings"
)

type simStats struct {
	dc int
	runs int
	hits int
	rerolls int
}

func newStat(dc int) *simStats {
	return &simStats{dc, 0, 0, 0}
}

func (ss *simStats) fraction() float64 {
	return (float64(ss.hits)/float64(ss.runs))
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
	result := rand.Intn(20) + 1 // roll d20
	if newRules && result <= s.mod {
		stats.rerolls++
		result = rand.Intn(20) + 1 // re-roll d20
	}
	if result == 20 { // a 20 always hits
		stats.hits++
	} else if result != 1 && result + s.mod >= stats.dc { // 1 always fails (not critical)
		stats.hits++
	}
	stats.runs++
}

func (s *sim) RunOnce() {
	for _, simStat := range s.oldStats {
		s.run(false, simStat)
	}
	for _, simStat := range s.newStats {
		s.run(true, simStat)
	}
}

func (s *sim) PrintStats() {
	fmt.Printf("Ability mod: +%d\n", s.mod)
	fmt.Printf("DC\tStandard\tHouse\tHouse rerolls\n")
	for i := 0; i < s.attempts; i++ {
		oldStat := s.oldStats[i]
		newStat := s.newStats[i]
		fmt.Printf("%d\t%.2f\t%.2f\t%d\n", oldStat.dc, oldStat.fraction()*100.0, newStat.fraction()*100.0, newStat.rerolls)
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

func (c *simContainer) RunAllOnce() {
	for _, sim := range c.sims {
		sim.RunOnce()
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
	rand.Seed(time.Now().UnixNano())

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

	max := 100000

	for i := 0; i < max; i++ {
		container.RunAllOnce()
	}

	container.PrintAllStats()
	container.DrawGraph()
}