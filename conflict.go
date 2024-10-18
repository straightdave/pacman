package main

type Conflict interface {
	Name() string
	Pos() []int
	OnConflicted(t Conflict)
}
