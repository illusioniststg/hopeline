package domain

type Survivor struct {
	FirstName string     `json:"first_name" db:"first_name"`
	LastName  string     `json:"last_name" db:"last_name"`
	DOB       string     `json:"dob" db:"dob"`
	Gender    string     `json:"gender" db:"gender"`
	Phone     string     `json:"phone" db:"phone"`
	Locs      []Location `json:"locs" db:"locs"`
	Village   string     `json:"village" db:"village"`
	District  string     `json:"district" db:"district"`
	State     string     `json:"state" db:"state"`
	IsInjured bool       `json:"is_injured" db:"is_injured"`
	IsMissing bool       `json:"is_missing" db:"is_missing"`
	ShelterID int64      `json:"shelter_id" db:"shelter_id"`
	BoatID    int64      `json:"boat_id" db:"boat_id"`
}

type SurvivorSearchRequest struct {
	Phone string `json:"phone"`
}
