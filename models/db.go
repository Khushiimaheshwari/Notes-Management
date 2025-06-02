package models

import (
    "fmt"
    "os"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASS"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"))
    
    fmt.Println("DSN:", dsn)

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    
    if err != nil {
        panic("Failed to connect to DB!")
        fmt.Println("Failed to connect to DB! Error:", err)
        panic(err)
    }

    db.AutoMigrate(&User{}, &Note{})
    DB = db
}
