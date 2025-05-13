package models

import "time"

type ActivityLog struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    UserID    uint      `json:"user_id" gorm:"column:user_id"`
    Accion    string    `json:"accion" gorm:"column:accion"`
    Detalles  string    `json:"detalles" gorm:"column:detalles"`  // Antes: Detalle
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` // Antes: FechaHora
}