package main

import (
	"fmt"
)

// Ana Product struct
type Product struct {
	Name  string
	Price float64
}

// Product interface
type ProductInterface interface {
	GetName() string
	GetPrice() float64
	CalculateShipping() float64
}

// Electronics struct (Product'u embed ediyor)
type Electronics struct {
	Product
}

func (e Electronics) GetName() string {
	return e.Name
}

func (e Electronics) GetPrice() float64 {
	return e.Price
}

func (e Electronics) CalculateShipping() float64 {
	return 20.0 // Elektronik için kargo
}

// Furniture struct (Product'u embed ediyor)
type Furniture struct {
	Product
}

func (f Furniture) GetName() string {
	return f.Name
}

func (f Furniture) GetPrice() float64 {
	return f.Price
}

func (f Furniture) CalculateShipping() float64 {
	return 50.0 // Mobilya için kargo
}

// Clothing struct (Product'u embed ediyor)
type Clothing struct {
	Product
}

func (c Clothing) GetName() string {
	return c.Name
}

func (c Clothing) GetPrice() float64 {
	return c.Price
}

func (c Clothing) CalculateShipping() float64 {
	return 5.0 // Kıyafet için kargo
}

// Sepet hesaplama fonksiyonu
func CalculateCart(cart []ProductInterface) (totalPrice, totalShipping float64) {
	for _, item := range cart {
		totalPrice += item.GetPrice()
		totalShipping += item.CalculateShipping()
		fmt.Printf("%s - Fiyat: %.2f, Kargo: %.2f\n", item.GetName(), item.GetPrice(), item.CalculateShipping())
	}
	return
}

func main() {
	cart := []ProductInterface{
		Electronics{Product{"Kulaklık", 150}},
		Furniture{Product{"Sandalye", 300}},
		Clothing{Product{"Tişört", 50}},
	}

	totalPrice, totalShipping := CalculateCart(cart)

	fmt.Printf("\nToplam Fiyat: %.2f\n", totalPrice)
	fmt.Printf("Toplam Kargo: %.2f\n", totalShipping)
	fmt.Printf("Genel Toplam: %.2f\n", totalPrice+totalShipping)
}
