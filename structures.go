package main

import "time"

type note struct {
	ID		  uint
	ChatID    int64
	Title     string
	FullText  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type todo struct {
	ID        uint
	ChatID    int64
	Task      string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
