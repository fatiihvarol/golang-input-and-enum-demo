package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Discount struct {
	Amount    float64
	IsPercent bool
}

type Config struct {
	FreeShippingThreshold float64
	KeyboardInputKey      string
	DiscountCodes         map[string]Discount
}

// Global config değişkeni
var config = Config{
	FreeShippingThreshold: 500.0,
	KeyboardInputKey:      "1",
	DiscountCodes: map[string]Discount{
		"DISC50":  {50, false},
		"DISC100": {100, false},
		"PERC5":   {0.05, true},
		"PERC10":  {0.30, true},
	},
}

type Product struct {
	Name         string
	Price        float64
	Stock        int
	DiscountRate float64
	ShippingCost float64
}

type ProductInterface interface {
	GetName() string
	GetPrice() float64
	GetStock() int
	CalculateShipping() float64
	CalculateDiscount() int
	ReduceStock(quantity int) bool
	ApplyExtraDiscount(rate float64)
}

type Electronics struct{ Product }

func (e Electronics) GetName() string            { return e.Name }
func (e Electronics) GetPrice() float64          { return e.Price }
func (e Electronics) GetStock() int              { return e.Stock }
func (e Electronics) CalculateShipping() float64 { return e.ShippingCost }
func (e Electronics) CalculateDiscount() int     { return int(e.Price * e.DiscountRate) }
func (e *Electronics) ReduceStock(quantity int) bool {
	if e.Stock >= quantity {
		e.Stock -= quantity
		return true
	}
	return false
}
func (e *Electronics) ApplyExtraDiscount(rate float64) { e.DiscountRate += rate }

type Furniture struct{ Product }

func (f Furniture) GetName() string            { return f.Name }
func (f Furniture) GetPrice() float64          { return f.Price }
func (f Furniture) GetStock() int              { return f.Stock }
func (f Furniture) CalculateShipping() float64 { return f.ShippingCost }
func (f Furniture) CalculateDiscount() int     { return int(f.Price * f.DiscountRate) }
func (f *Furniture) ReduceStock(quantity int) bool {
	if f.Stock >= quantity {
		f.Stock -= quantity
		return true
	}
	return false
}
func (f *Furniture) ApplyExtraDiscount(rate float64) { f.DiscountRate += rate }

type Clothing struct{ Product }

func (c Clothing) GetName() string            { return c.Name }
func (c Clothing) GetPrice() float64          { return c.Price }
func (c Clothing) GetStock() int              { return c.Stock }
func (c Clothing) CalculateShipping() float64 { return c.ShippingCost }
func (c Clothing) CalculateDiscount() int     { return int(c.Price * c.DiscountRate) }
func (c *Clothing) ReduceStock(quantity int) bool {
	if c.Stock >= quantity {
		c.Stock -= quantity
		return true
	}
	return false
}
func (c *Clothing) ApplyExtraDiscount(rate float64) { c.DiscountRate += rate }

type CartItem struct {
	Product  ProductInterface
	Quantity int
}

type Cart struct{ Items []CartItem }

func (c *Cart) AddItem(p ProductInterface, quantity int) bool {
	if p.GetStock() < quantity {
		fmt.Printf("Üzgünüz, %s stokta yeterli değil.\n", p.GetName())
		return false
	}
	p.ReduceStock(quantity)
	c.Items = append(c.Items, CartItem{Product: p, Quantity: quantity})
	return true
}

func (c *Cart) ApplyDiscountCode(d Discount) {
	if d.IsPercent {
		for _, item := range c.Items {
			item.Product.ApplyExtraDiscount(d.Amount)
		}
	} else {
		totalPrice := 0.0
		for _, item := range c.Items {
			totalPrice += item.Product.GetPrice() * float64(item.Quantity)
		}
		if totalPrice == 0 {
			return
		}
		for _, item := range c.Items {
			itemPrice := item.Product.GetPrice() * float64(item.Quantity)
			proportion := itemPrice / totalPrice
			extraRate := d.Amount * proportion / itemPrice
			item.Product.ApplyExtraDiscount(extraRate)
		}
	}
}

func (c Cart) CalculateTotals() (totalPrice, totalDiscount, totalShipping, grandTotal float64, freeShipping bool) {
	for _, item := range c.Items {
		price := item.Product.GetPrice() * float64(item.Quantity)
		discount := float64(item.Product.CalculateDiscount() * item.Quantity)
		shipping := item.Product.CalculateShipping() * float64(item.Quantity)
		totalPrice += price
		totalDiscount += discount
		totalShipping += shipping
	}

	grandTotal = totalPrice - totalDiscount
	if grandTotal >= config.FreeShippingThreshold {
		totalShipping = 0
		freeShipping = true
	}
	grandTotal += totalShipping
	return
}

func (c Cart) PrintCart() {
	fmt.Println("==== SEPET DETAYLARI ====")
	for _, item := range c.Items {
		price := item.Product.GetPrice() * float64(item.Quantity)
		discount := float64(item.Product.CalculateDiscount() * item.Quantity)
		shipping := item.Product.CalculateShipping() * float64(item.Quantity)
		finalPrice := price - discount + shipping
		fmt.Printf("%s x%d\n", item.Product.GetName(), item.Quantity)
		fmt.Printf("  Birim Fiyat: %.2f\n", item.Product.GetPrice())
		fmt.Printf("  Toplam Fiyat: %.2f\n", price)
		fmt.Printf("  Toplam İndirim: %.2f\n", discount)
		fmt.Printf("  Kargo: %.2f\n", shipping)
		fmt.Printf("  Ödenecek Tutar: %.2f\n\n", finalPrice)
	}
	totalPrice, totalDiscount, totalShipping, grandTotal, freeShipping := c.CalculateTotals()
	fmt.Printf("Sepet Toplam Fiyat: %.2f\n", totalPrice)
	fmt.Printf("Sepet Toplam İndirim: %.2f\n", totalDiscount)
	fmt.Printf("İndirimden sonraki Fiyat: %.2f\n", totalPrice-totalDiscount)
	if freeShipping {
		fmt.Println("Tebrikler! Kargo ücretsiz.")
	} else {
		fmt.Printf("Sepet Toplam Kargo: %.2f\n", totalShipping)
	}
	fmt.Printf("Sepet Genel Toplam: %.2f\n", grandTotal)
	fmt.Println("==========================")
}

func main() {
	headphones := &Electronics{Product{"Kulaklık", 150, 10, 0.10, 20}}
	chair := &Furniture{Product{"Sandalye", 300, 5, 0.15, 50}}
	tshirt := &Clothing{Product{"Tişört", 50, 20, 0.10, 5}}

	cart := Cart{}
	cart.AddItem(headphones, 2)
	cart.AddItem(chair, 1)
	cart.AddItem(tshirt, 3)

	cart.PrintCart()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("İndirim kodunuzu girin (çıkmak için Enter'a basın): ")
		codeInput, _ := reader.ReadString('\n')
		codeInput = strings.TrimSpace(strings.ToUpper(codeInput))

		if codeInput == "" {
			break // Enter basıldı, döngüden çık
		}

		if d, ok := config.DiscountCodes[codeInput]; ok {
			cart.ApplyDiscountCode(d)
			fmt.Println("İndirim kodu uygulandı!")
			cart.PrintCart()
			break
		} else {
			fmt.Printf("Geçersiz indirim kodu: '%s'. Tekrar deneyin.\n", codeInput)
		}
	}

}
