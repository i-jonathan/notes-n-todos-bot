package main

import (
	"fmt"
	"log"
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
	var notes []note
	db.Order("id").Where("chat_id = ?", chatID).Find(&notes)

	if len(notes) <= 1 {
		return "You have no notes."
	}

	message := "<b>Your Note(s):</b>\n"
	count := 0 

	for _, note := range notes {
		count ++
		message += fmt.Sprintf("%d. %s.\n", count, note.Title)
	}

	return message
}