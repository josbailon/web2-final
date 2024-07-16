package models/Payment

import (
    "gorm.io/gorm"
    "time"
)

type Payment struct {
    ID       uint           `gorm:"primaryKey"`
    Cantidad float64        `gorm:"type:numeric;not null"`
    Fecha    time.Time      `gorm:"type:timestamp;not null"`
    CreatedAt time.Time      `gorm:"autoCreateTime"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime"`
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
