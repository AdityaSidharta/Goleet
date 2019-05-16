package leetcode

import (
	"errors"
	"sort"
)

type coor struct {
	i int
	j int
}

type queue struct {
	values []coor
}

func (c1 coor) equal(c2 coor) bool {
	if (c1.i == c2.i) && (c1.j == c2.j) {
		return true
	} else {
		return false
	}
}

func (q *queue) put(c coor) {
	nvalues := len(q.values)
	for idx := 0; idx < nvalues; idx++ {
		if q.values[idx].equal(c) {
			return
		}
	}
	q.values = append(q.values, c)
}

func (q *queue) pop() (coor, error) {
	if len(q.values) == 0 {
		return coor{0, 0}, errors.New("no values within the queue")
	} else {
		result := q.values[0]
		q.values = q.values[1:]
		return result, nil
	}
}

func getCoor(forest [][]int) []coor {
	ni := len(forest)
	nj := len(forest[0])
	maps := make(map[int]coor)
	for i := 0; i < ni; i++ {
		for j := 0; j < nj; j++ {
			if forest[i][j] > 1 {
				maps[forest[i][j]] = coor{i, j}
			}
		}
	}

	keys := make([]int, 0, len(maps))
	values := make([]coor, 0, len(maps))

	for k := range maps {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		values = append(values, maps[k])
	}

	return values
}

func getNeighbour(forest [][]int, cur coor, maps map[coor]coor) []coor {
	ni := len(forest)
	nj := len(forest[0])
	result := make([]coor, 0)
	nxt := coor{cur.i + 1, cur.j + 1}
	prv := coor{cur.i - 1, cur.j - 1}
	if prv.j >= 0 {
		if forest[cur.i][prv.j] != 0 {
			if !isin(maps, coor{cur.i, prv.j}) {
				result = append(result, coor{cur.i, prv.j})
			}
		}
	}
	if nxt.j < nj {
		if forest[cur.i][nxt.j] != 0 {
			if !isin(maps, coor{cur.i, nxt.j}) {
				result = append(result, coor{cur.i, nxt.j})
			}
		}
	}
	if prv.i >= 0 {
		if forest[prv.i][cur.j] != 0 {
			if !isin(maps, coor{prv.i, cur.j}) {
				result = append(result, coor{prv.i, cur.j})
			}
		}
	}
	if nxt.i < ni {
		if forest[nxt.i][cur.j] != 0 {
			if !isin(maps, coor{nxt.i, cur.j}) {
				result = append(result, coor{nxt.i, cur.j})
			}
		}
	}
	return result
}

func isin(maps map[coor]coor, key coor) bool {
	_, ok := maps[key]
	return ok
}

func walk(forest [][]int, src coor, dst coor) int {
	if src.equal(dst) {
		return 0
	}
	maps := make(map[coor]coor)
	var q queue
	q.put(src)

	for len(q.values) != 0 {
		cur, _ := q.pop()
		for _, neigh := range getNeighbour(forest, cur, maps) {
			maps[neigh] = cur
			q.put(neigh)
		}
		if isin(maps, dst) {
			break
		}
	}
	if !isin(maps, dst) {
		return -1
	} else {
		result := 1
		cur := maps[dst]
		for !src.equal(cur) {
			cur = maps[cur]
			result = result + 1
		}
		return result
	}
}

func cutOffTree(forest [][]int) int {
	coors := getCoor(forest)
	coors = append([]coor{{0, 0}}, coors...)
	result := 0
	for idxCur := 0; idxCur < len(coors)-1; idxCur++ {
		idxNxt := idxCur + 1
		tmp := walk(forest, coors[idxCur], coors[idxNxt])
		if tmp == -1 {
			return -1
		} else {
			result = result + tmp
		}
	}
	return result
}
