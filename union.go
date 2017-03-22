// Package union provides an implementation of a union-find data structure
// derived from the work done by Robert Sedgewick and Kevin Wayne as part of
// Algorithms, 4th Edition.
//
// http://algs4.cs.princeton.edu/code/
//
// Copyright (C) 2000–2016 Robert Sedgewick and Kevin Wayne
// Copyright (C) 2017 Kasper Kronborg Isager
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more
// details.
//
// You should have received a copy of the GNU General Public License along with
// this program.  If not, see <http://www.gnu.org/licenses/>.
package union

// Union describes a union-find data structure used for partitioning a set of
// elements into groups of connected components.
type Union interface {
	// Find finds the parent of a given element.
	Find(p int) int

	// Join adds a connection between two elements.
	Join(p int, q int)

	// Connected checks if two elements are connected.
	Connected(p int, q int) bool
}

type union struct {
	parent map[int]int
	rank   map[int]int
}

// New constructs and initialises a new union-find data structure.
func New() Union {
	return &union{
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}
}

func (union union) Find(p int) int {
	root := p

	// Find the root of the element by following parent pointers until an element
	// without a parent is found.
	for {
		if q, ok := union.parent[root]; ok {
			root = q
		} else {
			break
		}
	}

	// Compress the connections between the element and the located root by making
	// every element found on the way to root point directly to it.
	for p != root {
		q := union.parent[p]
		union.parent[p] = root
		p = q
	}

	return root
}

func (union *union) Join(p int, q int) {
	pr := union.Find(p)
	qr := union.Find(q)

	if pr == qr {
		return
	}

	// Merge the lower-ranking component into the larger-ranking component.
	if union.rank[pr] < union.rank[qr] {
		union.parent[pr] = qr
	} else {
		union.parent[qr] = pr
	}

	// Increase the rank of the merged component if joining two components that
	// have the same rank.
	if union.rank[pr] == union.rank[qr] {
		union.rank[pr]++
	}
}

func (union union) Connected(p int, q int) bool {
	return union.Find(p) == union.Find(q)
}
