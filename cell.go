package main

// Cell - Representation of a cell
type Cell struct {
	State bool
}

// NewCell - returns a new cell
func NewCell(state bool) *Cell {
	return &Cell{State: state}
}
