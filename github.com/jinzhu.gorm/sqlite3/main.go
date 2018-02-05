package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "fmt"
)

type Product struct {
    gorm.Model
    Code string
    Price uint
}

func main() {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&Product{})

    // Create
    db.Create(&Product{Code: "L1212", Price: 1000})
    db.Create(&Product{Code: "ABC", Price: 1001})

    // Read
    var product Product
    id := 1
    db.First(&product, id) // find product with id 1
    fmt.Printf("product=%v\n", product)
    db.First(&product, "code = ?", "L123") // find product with code l1212
    fmt.Printf("product=%v\n", product)

    // Update - update product's price to 2000
    db.Model(&product).Update("Price", 2000)
    db.First(&product, "code = ?", "L123") // find product with code l1212
    fmt.Printf("product=%v\n", product)

    // Delete - delete product
    db.Delete(&product)
}