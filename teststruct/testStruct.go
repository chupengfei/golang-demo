package main
import (
	"encoding/json"
	"fmt"
)

// Product 商品信息
type Product struct {
	Name      string
	ProductID int64
	Number    int
	Price     float64
	IsOnSale  bool
}

func main() {
	p := &Product{}
	p.Name = "Xiao mi 6"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 1
	data, _ := json.Marshal(p)
	fmt.Println(p)
	fmt.Println(string(data))

	b := &Product{}
	b.Name = "Xiao mi 8"
	b.IsOnSale = true
	b.Number = 2000
	b.Price = 3499.00
	b.ProductID = 2
	datab, _ := json.Marshal(b)
	fmt.Println(b)
	fmt.Println(string(datab))
	fmt.Println(p)

}