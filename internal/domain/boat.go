package domain

type Boat struct {
    ID           string `json:"id" db:"id"`
    Name         string `json:"name" db:"name"`
    Capacity     int64  `json:"capacity" db:"capacity"`
    Availability int64  `json:"availability" db:"availability"`
    UpdatedAt    string `json:"updated_at" db:"updated_at"`
}