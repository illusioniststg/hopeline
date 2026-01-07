package domain

type Survivor struct {
    ID          int64      `json:"id" db:"id"`
    FirstName   string     `json:"first_name" db:"first_name"`
    LastName    string     `json:"last_name" db:"last_name"`
    DOB         string     `json:"dob" db:"dob"`
    Phone       string     `json:"phone" db:"phone"`
    Locs        []Location `json:"locs" db:"locs"`
    IsInjured   bool       `json:"is_injured" db:"is_injured"`
    IsMissing   bool       `json:"is_missing" db:"is_missing"`
    ShelterID   int64      `json:"shelter_id" db:"shelter_id"`
    BoatID      int64      `json:"boat_id" db:"boat_id"`
    UpdatedAt   string     `json:"updated_at" db:"updated_at"`
}