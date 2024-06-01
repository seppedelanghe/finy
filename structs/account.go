package structs

type Account struct  {
	BaseModel

	Name string
	Type string // TODO: update to type
	Balance int // in cents
	Currency string // TODO: map to currency type

	UserId string
}
