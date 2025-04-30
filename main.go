package main

import (
	"fmt"
	"github.com/JasperMunene/notes-api/models"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	n := models.Note{
		ID:        1,
		Title:     "Test",
		Content:   "Example",
		CreatedAt: time.Now(),
	}
	fmt.Println(n)
}
