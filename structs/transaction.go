package structs


type Transaction struct {
	BaseModel

	Name string `json:"name"`
	Amount int `json:"amount"` // in cents
	Date string `json:"date"`

	Origin *Account
	Destination *Account
}

