package models

import "time"

type JobStatus string

const (
	StatusPending    JobStatus = "pending"
	StatusProcessing JobStatus = "processing"
	StatusCompleted  JobStatus = "completed"
	StatusFailed     JobStatus = "failed"
)

type Job struct {
	ID         string    `gorm:"primaryKey"`
	Status     JobStatus `gorm:"type:varchar(20);not null"`
	Payload    string
	Retries    int     `gorm:"default:0"`
	MaxRetries int     `gorm:"default:3"`
	ErrorMsg   *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
