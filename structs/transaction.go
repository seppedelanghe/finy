package structs


type Transaction struct {
	BaseModel

	Name string
	Amount int // in cents

	Origin *Account
	Destination *Account
}

