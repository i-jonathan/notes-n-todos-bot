package main

import (
	"log"
	"strings"
)

func processRequest(body *webHookReqBody) {
	parts := strings.Fields(body.Message.Text)
	userID := body.Message.Chat.ID

	var command string
	if len(parts) > 0 {
		if strings.HasSuffix(parts[0], "notes_n_todos_bot") {
			command = strings.Split(parts[0], "@")[0]
		} else {
			command = parts[0]
		}
	}

	switch len(parts) {
	case 1:
		switch command {
		case "/help":
			helpText := "/help - Display help text.\n\n<b>Todo Commands:</b\n" +
				"/addtask task-name - Creates a todo item with the indicated name.\n\n" +
				"/donetask number(s) - Marks indicated Todo items as done." +
				"Use the number displayed from /viewtodolist. For multiple numbers, separate them with a space.\n\n" +
				"/viewtodolist - List all your items on your Todo list.\n"

			if err := respond(userID, helpText); err != nil {
				log.Println(err)
			}
		case "/viewtodolist":
			text := listTodos(db, userID)
			err := respond(userID, text)
			if err != nil {
				log.Println(err)
			}
		case "/cleantodolist":
			done := cleanTodos(db, userID)
			if done {
				if err := respond(userID, "Your todo list has been cleaned."); err != nil {
					log.Println(err)
				}
			} else {
				if err := respond(userID, "An error occurred. Try again later."); err != nil {
					log.Println(err)
				}
			}
		default:
			err := respond(userID, "Hi.\nType /help to get the help text")
			if err != nil {
				log.Println(err)
			}
			return
		}
	// case 2:

	default:
		switch command {
		case "/donetask":
			var numbers []string
			numbers = append(numbers, parts[1:]...)
			result := markTodo(db, userID, numbers)

			if err := respond(userID, result); err != nil {
				log.Println(err)
			}
		case "/addtask":
			text := strings.Join(parts[1:], " ")
			isSuccess, err := createTodo(db, text, userID)
			if isSuccess == false {
				log.Println(err)
				err := respond(userID, "An error occurred. \nYour Todo item wasn't added.")
				if err != nil {
					log.Println(err)
				}
				return
			}

			err = respond(userID, "Your Todo Item has been created successfully.")
			if err != nil {
				log.Println(err)
			}
			return
		default:
			helpText := "/help - Display help text.\n\n<b>Todo Commands:</b\n" +
				"/addtask task-name - Creates a todo item with the indicated name.\n\n" +
				"/donetask number(s) - Marks indicated Todo items as done." +
				"Use the number displayed from /viewtodolist. For multiple numbers, separate them with a space.\n\n" +
				"/viewtodolist - List all your items on your Todo list.\n"

			if err := respond(userID, helpText); err != nil {
				log.Println(err)
			}
		}
	}
}
