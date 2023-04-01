package product

type ID uint32

type Product struct {
	ID          ID      // 商品ID
	Name        string  // 商品名
	Description string  // 商品説明
	Price       float64 // 商品価格
}
