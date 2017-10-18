package main

import (
	"fmt"
	"strings"
	"math"
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
	botStats []*simStats
	mod int
	startDc int
	attempts int
}

func newSim(mod int, basedc int, dctries int) *sim {
	var oldStats []*simStats
	var newStats []*simStats
	var botStats []*simStats
	for i := 0; i < dctries; i++ {
		oldStats = append(oldStats, newStat(basedc+i))
		newStats = append(newStats, newStat(basedc+i))
		botStats = append(botStats, newStat(basedc+i))
	}
	return &sim{
		oldStats,
		newStats,
		botStats,
		mod,
		basedc,
		dctries,
	}
}

func (s *sim) run(newRules bool, stats *simStats) {
	// TODO: Rewrite this to use the formula for the whole thing.
	mod := s.mod
	dc := stats.dc
	for i := 1; i <= 20; i++ {
		if newRules && i <= mod {
			for j := 1; j <= 20; j++ {
				if j == 20 {
					stats.hits += 0.05
				} else if j != 1 && j + mod >= dc { // 1 always fails (not critical)
					stats.hits += 0.05
				}
			}
		} else {
			if i == 20 {
				stats.hits++
			} else if i != 1 && i + mod >= dc { // 1 always fails (not critical)
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
	for _, simStat := range s.botStats {
		//simStat.hits = ((21.0 - float64(simStat.dc) + float64(s.mod))/20.0) * ((float64(s.mod) + 20.0)/20.0)
		m := math.Max((float64(simStat.dc) - 1.0 - float64(s.mod))/20.0, 0)
		simStat.hits = math.Min(1.0, (1.0-m) + (1-m)*math.Min(m, math.Ceil(m)*float64(s.mod)/20.0))
		//simStat.hits = ((float64(s.mod)-1.0)/20.0)*((21.0-float64(simStat.dc)+float64(s.mod))/20.0)
		simStat.runs = 1
	}
}

func (s *sim) PrintStats() {
	fmt.Printf("Ability mod: +%d\n", s.mod)
	fmt.Printf("DC\tStandard\tHouse\tbot\n")
	for i := 0; i < s.attempts; i++ {
		oldStat := s.oldStats[i]
		newStat := s.newStats[i]
		botStat := s.botStats[i]
		fmt.Printf("%d\t%.4f\t%.4f\t%.4f\n", oldStat.dc, oldStat.fraction(), 
			newStat.fraction(), botStat.fraction())
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
		var botPoints []point
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
		for _, stat := range sim.botStats {
			x := (stat.dc-sim.startDc) * pointWidth
			y := height - int(stat.fraction()*float64(height))
			botPoints = append(botPoints, point{x,y})
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
				foundPoint := []bool{false, false, false}
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
				for _, p := range botPoints {
					if p.x == w && p.y == h {
						foundPoint[2] = true
					}
				}
				if foundPoint[0] && foundPoint[1] {
					fmt.Printf(" @")
				} else if foundPoint[1] && foundPoint[2] {
					fmt.Printf(" &")
				} else if foundPoint[0] {
					fmt.Printf(" *")
				} else if foundPoint[1] {
					fmt.Printf(" #")
				} else if foundPoint[2] {
					fmt.Printf(" $")
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
