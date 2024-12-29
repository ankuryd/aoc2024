package day23

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"aoc2024/util"
)

func solve1(graph map[string]map[string]struct{}) string {
	triplets := make(map[string]struct{})
	for u, neighbors := range graph {
		for v := range neighbors {
			if v <= u {
				continue
			}

			for w := range neighbors {
				if w <= u || w == v {
					continue
				}

				if _, ok := graph[v][w]; ok {
					nodes := []string{u, v, w}
					slices.Sort(nodes)
					triplets[strings.Join(nodes, "")] = struct{}{}
				}
			}
		}
	}

	result := 0
	for triplet := range triplets {
		if triplet[0] == 't' || triplet[2] == 't' || triplet[4] == 't' {
			result++
		}
	}

	return fmt.Sprintf("%d", result)
}

func solve2(graph map[string]map[string]struct{}) string {
	triplets := make(map[string]map[string]struct{})
	for u, neighbors := range graph {
		for v := range neighbors {
			if v <= u {
				continue
			}

			for w := range neighbors {
				if w <= u || w == v {
					continue
				}

				if _, ok := graph[v][w]; ok {
					if _, ok := triplets[u]; !ok {
						triplets[u] = make(map[string]struct{})
					}
					triplets[u][v] = struct{}{}
					triplets[u][w] = struct{}{}

					if _, ok := triplets[v]; !ok {
						triplets[v] = make(map[string]struct{})
					}
					triplets[v][u] = struct{}{}
					triplets[v][w] = struct{}{}

					if _, ok := triplets[w]; !ok {
						triplets[w] = make(map[string]struct{})
					}
					triplets[w][u] = struct{}{}
					triplets[w][v] = struct{}{}
				}
			}
		}
	}

	cliques := make(map[string]int)
	for k, v := range triplets {
		nodes := make([]string, 0)
		nodes = append(nodes, k)
		for k2 := range v {
			nodes = append(nodes, k2)
		}
		slices.Sort(nodes)
		cliques[strings.Join(nodes, ",")]++
	}

	result := ""
	for clique, count := range cliques {
		// 2*count + count - 1 = len(clique) because each node have 2 characters and we have count-1 commas
		if 3*count-1 == len(clique) && len(clique) > len(result) {
			result = clique
		}
	}

	return result
}

func Run(day int, input []string) {
	graph := make(map[string]map[string]struct{})

	for i, line := range input {
		if line == "" {
			util.Fatal("Invalid format on line %d: empty line", i)
		}

		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			util.Fatal("Invalid format on line %d: expected two parts", i)
		}

		if _, ok := graph[parts[0]]; !ok {
			graph[parts[0]] = make(map[string]struct{})
		}

		if _, ok := graph[parts[1]]; !ok {
			graph[parts[1]] = make(map[string]struct{})
		}

		graph[parts[0]][parts[1]] = struct{}{}
		graph[parts[1]][parts[0]] = struct{}{}
	}

	startTime := time.Now()
	util.Output(1, solve1(graph))
	elapsed := time.Since(startTime)
	util.TimeTaken(elapsed)

	util.Separator()

	startTime = time.Now()
	util.Output(2, solve2(graph))
	elapsed = time.Since(startTime)
	util.TimeTaken(elapsed)
}
