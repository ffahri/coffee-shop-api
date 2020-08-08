package v1

type DrinkType string

const (
	TEA    DrinkType = "TEA"
	COFFEE DrinkType = "COFFEE"
	SODA   DrinkType = "SODA"
)

type drinks struct {
	Id        string    `db:"Id"`
	Name      string    `db:"Name"`
	Type      DrinkType `db:"Type"`
	Origin    string    `db:"Origin"`
	Brand     string    `db:"Brand"`
	Price     float64   `db:"Price"`
	Stock     int       `db:"Stock"`
	Timestamp []uint8   `db:"Timestamp"`
}

type OrderRequest struct {
	Id       string
	Quantity int
}
