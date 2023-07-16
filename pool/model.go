package pool

import "time"

type Status int

const (
	StatusCreated Status = 0
	StatusRunning Status = 1
	StatusSuccess Status = 2
	StatusFail    Status = 3
)

type Job struct {
	ID        int64     `gorm:"primaryKey,column:id"`
	Host      string    `gorm:"column:host"`
	Status    Status    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Job) TableName() string {
	return "job"
}
