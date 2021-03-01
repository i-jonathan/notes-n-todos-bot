package main

import (
	"log"
	"strings"
)

type messageStream struct {
	ChatID	int64
	Messages []string
}

var messageRegistery []*messageStream

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
						"/addtask <pre>task-name</pre> - Creates a todo item with the indicated name.\n\n" +
						"/donetask <pre>number(s)</pre> - Marks indicated Todo items as done." +
						"Use the number displayed from /viewtodolist. For multiple numbers, separate them with a space.\n\n" +
						"/viewtodolist - List all your items on your Todo list.\n\n\n" +
						"<b>Note Commands:</b>\n/addnote <pre>title</pre> <pre>note content</pre>  - Adds a new note with title and content.\n\n" +
						"/deletenote <pre>1 5</pre> - Deletes notes 1 and 5 (an example)\n" + 
						"/listnotes - Lists all available notes.\n" +
						"/readnote <pre>1</pre> - Reads note 1"

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
		case "/listnotes":
			text := listNotes(db, userID)
			err := respond(userID, text)
			if err != nil {
				log.Println(err)
			}
		default:
			err := respond(userID, "Hi.\nType /help to get the help text")
			if err != nil {
				log.Println(err)
			}
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
		case "/addnote":
			title := parts[1]
			main := strings.Join(parts[2:], " ")
			added := addNote(db, userID, title, main)

			if added {
				err := respond(userID, "Your Note has been created successfully.")
				if err != nil {
					log.Println(err)
				}
			} else {
				err := respond(userID, "An Error Occurred, please try again later.")
				if err != nil {
					log.Println(err)
				}
			}
		case "/deletenote":
			var numbers []string
			numbers = append(numbers, parts[1:]...)
			text := deleteNote(db, userID, numbers)
			err := respond(userID, text)
			if err != nil {
				log.Println(err)
			}
		case "/readnote":
			text := noteDetails(db, userID, parts[1])
			err := respond(userID, text)
			if err != nil {
				log.Println(err)
			}
		default:
			helpText := "/help - Display help text.\n\n<b>Todo Commands:</b\n" +
			"/addtask <pre>task-name</pre> - Creates a todo item with the indicated name.\n\n" +
			"/donetask <pre>number(s)</pre> - Marks indicated Todo items as done." +
			"Use the number displayed from /viewtodolist. For multiple numbers, separate them with a space.\n\n" +
			"/viewtodolist - List all your items on your Todo list.\n\n\n" +
			"<b>Note Commands:</b>\n/addnote <pre><title></pre> <pre><note content</pre>  - Adds a new note with title and content.\n\n" +
			"/deletenote <pre>1 5</pre> - Deletes notes 1 and 5 (an example)\n" + 
			"/listnotes - Lists all available notes.\n" +
			"/readnote <pre>1</pre> - Reads note 1"

			if err := respond(userID, helpText); err != nil {
				log.Println(err)
			}
		}
	}
}


func findElement(slice []*messageStream, chatID int64) (int, bool) {
	for i, item := range slice {
		if item.ChatID == chatID {
			return i, true
		}
	}

	return -1, false
}