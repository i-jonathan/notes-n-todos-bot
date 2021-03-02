package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

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

	if len(notes) == 0 {
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

func deleteNote(db *gorm.DB, chatID int64, noteIndexes []string) string {
	var notes []note
	db.Order("id").Where("chat_id = ?", chatID).Find(&notes)

	sort.Strings(noteIndexes)

	for i := len(noteIndexes); i > 0; i-- {
		number, _ := strconv.Atoi(noteIndexes[i-1])
		if number > len(notes) || number < 0 {
			return "Please input a valid number."
		}
		err := db.Where("id = ?", notes[number-1].ID).Delete(&note{}).Error
		if err != nil {
			log.Println(err)
			return "An error occured. Please try again later."
		}	
	}	
	return "Deleted Successfully"
}

func noteDetails(db *gorm.DB, chatID int64, noteIndex string) string {
	var notes []note
	db.Order("id").Where("chat_id = ?", chatID).Find(&notes)
	index, _ := strconv.Atoi(noteIndex)

	if index > len(notes) || index < 1 {
		return "Please input a valid number."
	}
	note := notes[index-1]
	text := fmt.Sprintf("<b>Title:</b> %s\n\n<b>Main Text:</b> %s", note.Title, note.FullText)

	return text
}