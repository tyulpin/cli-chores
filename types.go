package main

import "time"

type Chore struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Reminder    string    `json:"reminder"`
}

type ChoreManager struct {
	Chores []Chore `json:"chores"`
}
