package models

import "time"

type Vehicle struct {
    ID      uint      `gorm:"primaryKey" xml:"id"`
    Plate   string    `xml:"plate"`
    InTime  time.Time `xml:"in_time"`
    OutTime time.Time `xml:"out_time"`
}
