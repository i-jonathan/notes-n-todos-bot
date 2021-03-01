package main

import "time"

type note struct {
	chatID    string
	title     string
	full      string
	createdAt time.Time
	updatedAt time.Time
}

type todo struct {
	ID        uint
	ChatID    int64
	Task      string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
