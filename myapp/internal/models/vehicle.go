package models/Vehicle
import (
    "gorm.io/gorm"
    "time"
)

type Vehicle struct {
    ID     uint      `gorm:"primaryKey"`
    Marca  string    `gorm:"type:varchar(50);not null"`
    Modelo string    `gorm:"type:varchar(50);not null"`
    Año    int       `gorm:"type:int;not null"`
    CreatedAt time.Time      `gorm:"autoCreateTime"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime"`
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
