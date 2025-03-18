package main

import (
	"fmt"
	"io"
	"os"
	"slices"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(t1 string, s1 int, t2 string, s2 int) {
	if s1 > s2 {
		l.Wins[t1] += 1
	} else if s1 < s2 {
		l.Wins[t2] += 1
	}
}

func (l League) Ranking() []string {
	r := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		r = append(r, k)
	}
	sortF := func(t1 string, t2 string) int {
		return l.Wins[t2] - l.Wins[t1]
	}
	slices.SortFunc(r, sortF)
	return r
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	for _, v := range r.Ranking() {
		io.WriteString(w, v+"\n")
	}
}

func main() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)
	results := l.Ranking()
	fmt.Println(results)

	RankPrinter(l, os.Stdout)
}
