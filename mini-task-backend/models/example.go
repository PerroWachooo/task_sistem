package models

import (
    "time"
)

type Example struct {
    ID    string `json:"id" bson:"_id,omitempty"`
    Title string `json:"title" bson:"title"`
}

type Task struct {
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed"`
    CreatedAt   time.Time `json:"createdAt"`
}