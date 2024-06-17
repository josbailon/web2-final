package models

import "time"

type Payment struct {
    ID        uint      `gorm:"primaryKey" xml:"id"`
    VehicleID uint      `xml:"vehicle_id"`
    Amount    float64   `xml:"amount"`
    Time      time.Time `xml:"time"`
}
