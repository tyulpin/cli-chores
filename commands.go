package main

import (
	"fmt"
	"log"
	"time"
)

func addChore(manager *ChoreManager, description, dueDateStr, reminder string) {
	dueDate, err := time.Parse("2006-01-02", dueDateStr)
	if err != nil {
		log.Fatalf("Invalid date format. Use YYYY-MM-DD: %v", err)
	}
	id := len(manager.Chores) + 1
	chore := Chore{
		ID:          id,
		Description: description,
		DueDate:     dueDate,
		Reminder:    reminder,
	}
	manager.Chores = append(manager.Chores, chore)
	saveChores(*manager)
	fmt.Println("Chore added successfully!")
}

func listChores(manager ChoreManager) {
	if len(manager.Chores) == 0 {
		fmt.Println("No chores available.")
		return
	}
	fmt.Println("Current Chores:")
	for _, chore := range manager.Chores {
		fmt.Printf("ID: %d, Description: %s, Due Date: %s, Reminder: %s\n", chore.ID, chore.Description, chore.DueDate.Format("2006-01-02"), chore.Reminder)
	}
}

func deleteChore(manager *ChoreManager, id int) {
	index := -1
	for i, chore := range manager.Chores {
		if chore.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Chore not found.")
		return
	}
	manager.Chores = append(manager.Chores[:index], manager.Chores[index+1:]...)
	saveChores(*manager)
	fmt.Println("Chore deleted successfully!")
}
