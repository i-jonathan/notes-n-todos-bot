package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func createTodo(db *gorm.DB, name string, chatID int64) (bool, error) {

	if strings.ReplaceAll(name, " ", "") == "" {
		return false, errors.New("Empty task name")
	}

	newTodo := todo{
		ChatID: chatID,
		Task:   name,
		Done:   false,
	}

	result := db.Create(&newTodo)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func listTodos(db *gorm.DB, chatID int64) string {
	var todos []todo
	db.Order("id").Where("chat_id = ?", chatID).Find(&todos)

	if len(todos) == 0 {
		return "You have no Todos."
	}

	message := "Your To-do(s):\n"
	count := 0

	for _, item := range todos {
		count++
		if item.Done != true {
			message += fmt.Sprintf("%d. %s.\n", count, item.Task)
		} else {
			message += fmt.Sprintf("%d. <s>%s.</s>\n", count, item.Task)

		}
	}

	return message
}

func markTodo(db *gorm.DB, chatID int64, numbers []string) string {
	var todos []todo

	db.Order("id").Where("chat_id = ?", chatID).Find(&todos)
	if len(todos) == 0 {
		return "You have no Todos."
	}

	sort.Strings(numbers)

	for i := len(numbers); i > 0; i-- {
		number, _ := strconv.Atoi(numbers[i-1])
		if number > len(todos) || number < 1 {
			return "Please enter a valid number"
		}
		todo := todos[number-1]
		todo.Done = true
		db.Save(&todo)
	}

	return "Marked as Done."
}

func cleanTodos(db *gorm.DB, chatID int64) bool {
	err := db.Where("chat_id = ? AND done = ?", chatID, true).Delete(&todo{}).Error

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}