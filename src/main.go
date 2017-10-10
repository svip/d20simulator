package main

import (
	"math/rand"
	"time"
	"fmt"
)

var proficiencyTable = map[int]int {
	1: 2,
	2: 2,
	3: 2,
	4: 2,
	5: 3,
	6: 3,
	7: 3,
	8: 3,
	9: 4,
	10: 4,
	11: 4,
	12: 4,
	13: 5,
	14: 5,
	15: 5,
	16: 5,
	17: 6,
	18: 6,
	19: 6,
	20: 6,
}

type weapon struct {
	useStr bool
	mod int
}

func newWeapon(useStr bool, mod int) weapon {
	return weapon{
		useStr,
		mod,
	}
}

type character struct {
	str int
	dex int
	level int
}

func newChar(str, dex, level int) character {
	return character{
		str,
		dex,
		level,
	}
}

func (c *character) proficiency() int {
	p, ok := proficiencyTable[c.level]
	if !ok { // wrong level!
		c.level = 1
		p = proficiencyTable[c.level]
	}
	return p
}

func (c *character) weaponMod(w weapon) int {
	if w.useStr {
		return c.proficiency() + abilityToMod(c.str) + w.mod
	} else {
		return c.proficiency() + abilityToMod(c.dex) + w.mod
	}
}

type simStats struct {
	ac int
	runs int
	hits int
	rerolls int
}

func newStats(ac int) *simStats {
	return &simStats{ac, 0, 0, 0}
}

type sim struct {
	stats []*simStats
	w weapon
	c character
	newRules bool
}

func newSim(newRules bool, w weapon, c character, noo int) sim {
	statsList := []*simStats{}
	for i := 0; i < noo; i++ {
		statsList = append(statsList, newStats(10+i))
	}
	return sim{
		statsList,
		w,
		c,
		newRules,
	}
}

func abilityToMod(ability int) int {
	mod := ability - 10
	mod = mod/2
	return mod
}

func (s *sim) RunOnce() {
	for _, stat := range s.stats {
		result := rand.Intn(20) + 1 // roll d20
		bonus := s.c.weaponMod(s.w)
		if s.newRules && result <= bonus {
			stat.rerolls++
			result = rand.Intn(20) + 1 // re-roll d20
		}
		if result + bonus >= stat.ac {
			stat.hits++
		}
		stat.runs++
	}
}

func (s *sim) PrintStats(title string) {
	fmt.Printf("Character (level %d): str: %d (+%d), dex: %d (+%d), proficiency: %d\n", s.c.level, 
		s.c.str, abilityToMod(s.c.str), s.c.dex, abilityToMod(s.c.dex), s.c.proficiency())
	fmt.Printf("%s: +%d with mods: +%d\n", title, s.w.mod, s.c.weaponMod(s.w))
	for _, stat := range s.stats {
		fmt.Printf("AC: %d: Runs: %d, Hits: %d, Percentage: %.2f %% (re-rolls: %d)\n", stat.ac, stat.runs, stat.hits,
			(float64(stat.hits)/float64(stat.runs))*100.0, stat.rerolls)
	}
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	sword := newWeapon(true, 0)
	bow := newWeapon(false, 0)
	char := newChar(16, 12, 3)

	regularSword := newSim(false, sword, char, 12)
	regularBow := newSim(false, bow, char, 12)
	newSword := newSim(true, sword, char, 12)
	newBow := newSim(true, bow, char, 12)

	max := 10000

	for i := 0; i < max; i++ {
		regularSword.RunOnce()
		regularBow.RunOnce()
		newSword.RunOnce()
		newBow.RunOnce()
	}

	regularSword.PrintStats("regular sword")
	regularBow.PrintStats("regular bow")
	newSword.PrintStats("new sword")
	newBow.PrintStats("new bow")
}