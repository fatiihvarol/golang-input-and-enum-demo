package main

import (
	"fmt"
)

type Student struct {
	FirstName string
	LastName  string
	Age       int
	Grade     float32
}

// Pointer ile öğrenci notunu güncelleyen fonksiyon
func updateGrade(s *Student, newGrade float32) {
	s.Grade = newGrade
}

// Value ile öğrenci notunu güncelleyen fonksiyon (kopya üzerinden çalışır)
func updateGradeValue(s Student, newGrade float32) {
	s.Grade = newGrade
}

func main() {
	// Normal öğrenci oluşturuyoruz
	student := Student{
		FirstName: "Ahmet",
		LastName:  "Yılmaz",
		Age:       20,
		Grade:     3.5,
	}

	// Pointer ile öğrenci oluşturuyoruz (eskiyi BOZMADAN)
	studentPtr := &Student{
		FirstName: "Mehmet",
		LastName:  "Kaya",
		Age:       22,
		Grade:     3.8,
	}

	// Eski öğrenci bilgileri
	fmt.Println("---- Normal Öğrenci ----")
	fmt.Println("Ad:", student.FirstName)
	fmt.Println("Soyad:", student.LastName)
	fmt.Println("Yaş:", student.Age)
	fmt.Println("Not Ortalaması:", student.Grade)

	//Pointer Addresleri
	fmt.Printf("Normal Öğrencinin Adresi: %p\n", &student)
	fmt.Printf("Pointer Öğrencinin Adresi: %p\n", studentPtr)

	fmt.Printf("Pointer Öğrencinin Adresi: %p\n", &studentPtr.FirstName)
	fmt.Printf("Pointer Öğrencinin Adresi: %p\n", &studentPtr.LastName)
	fmt.Printf("Pointer Öğrencinin Adresi: %p\n", &studentPtr.Age)

	// Pointer öğrenci bilgileri
	fmt.Println("\n---- Pointer Öğrenci ----")
	fmt.Println("Ad:", studentPtr.FirstName)
	fmt.Println("Soyad:", studentPtr.LastName)
	fmt.Println("Yaş:", studentPtr.Age)
	fmt.Println("Not Ortalaması:", studentPtr.Grade)

	// Pointer üzerinden güncelleme
	studentPtr.Grade = 4.0
	fmt.Println("\nPointer Öğrencinin Güncellenmiş Ortalaması:", studentPtr.Grade)

	// Fonksiyon ile güncelleme (pointer)
	updateGrade(studentPtr, 4.5)
	fmt.Println("Pointer ile Fonksiyondan Sonra Güncellenmiş Ortalaması:", studentPtr.Grade) // Değişir

	// Fonksiyon ile güncelleme (value) - orijinal struct değişmez
	updateGradeValue(student, 4.8)
	fmt.Println("Value ile Fonksiyondan Sonra Normal Öğrenci Ortalaması:", student.Grade) // Değişmez
}
