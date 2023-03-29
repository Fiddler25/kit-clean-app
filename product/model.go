package product

type ID uint32

type Product struct {
	ID          ID
	Name        string
	Description string
	Price       float64
	Stock       uint8
}
