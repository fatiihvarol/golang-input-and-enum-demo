package main

import (
	"fmt"
)

// Product interface
type Product interface {
	GetName() string
	GetPrice() float64
	CalculateShipping() float64
}

// Electronics struct
type Electronics struct {
	Name  string
	Price float64
}

func (e Electronics) GetName() string {
	return e.Name
}

func (e Electronics) GetPrice() float64 {
	return e.Price
}

func (e Electronics) CalculateShipping() float64 {
	return 20.0 // Kırılabilir elektronik ürünler için kargo
}

// Furniture struct
type Furniture struct {
	Name  string
	Price float64
}

func (f Furniture) GetName() string {
	return f.Name
}

func (f Furniture) GetPrice() float64 {
	return f.Price
}

func (f Furniture) CalculateShipping() float64 {
	return 50.0 // Büyük mobilya ürünleri için kargo
}

// Clothing struct
type Clothing struct {
	Name  string
	Price float64
}

func (c Clothing) GetName() string {
	return c.Name
}

func (c Clothing) GetPrice() float64 {
	return c.Price
}

func (c Clothing) CalculateShipping() float64 {
	return 5.0 // Küçük kıyafet ürünleri için kargo
}

// Sepet fonksiyonu
func CalculateCart(cart []Product) (totalPrice, totalShipping float64) {
	for _, item := range cart {
		totalPrice += item.GetPrice()
		totalShipping += item.CalculateShipping()
		fmt.Printf("%s - Fiyat: %.2f, Kargo: %.2f\n", item.GetName(), item.GetPrice(), item.CalculateShipping())
	}
	return
}

func main() {
	cart := []Product{
		Electronics{Name: "Kulaklık", Price: 150},
		Furniture{Name: "Sandalye", Price: 300},
		Clothing{Name: "Tişört", Price: 50},
	}

	totalPrice, totalShipping := CalculateCart(cart)

	fmt.Printf("\nToplam Fiyat: %.2f\n", totalPrice)
	fmt.Printf("Toplam Kargo: %.2f\n", totalShipping)
	fmt.Printf("Genel Toplam: %.2f\n", totalPrice+totalShipping)
}
