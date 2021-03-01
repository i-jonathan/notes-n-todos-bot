package main

import (
	"log"
	"strings"

	"gorm.io/gorm"
)

func addNote(db *gorm.DB, chatID int64, title, mainText string) bool {

	newNote := &note{
		ChatID:	chatID,
		Title: title,
		FullText: mainText,
	}

	result := db.Create(newNote)
	if result.Error != nil {
		log.Println(result.Error)
		return false
	}

	return true
}

func listNotes(db *gorm.DB, chatID int64) string {
	
}