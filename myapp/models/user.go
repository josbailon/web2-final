package models

type User struct {
    ID       uint   `gorm:"primaryKey" xml:"id"`
    Username string `xml:"username"`
    Password string `xml:"password"`
}
