package domain

type Shelter struct {
	ID        string
	Locs      []Location
	Capacity  int
	Available int
	Type      string
}
