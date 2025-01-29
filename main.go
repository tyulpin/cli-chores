package main

import "fmt"

func main() {
	manager := loadChores()

	for {
		fmt.Println("\nChore Manager")
		fmt.Println("1. Add Chore")
		fmt.Println("2. List Chores")
		fmt.Println("3. Delete Chore")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var description, dueDateStr, reminder string
			fmt.Print("Enter chore description: ")
			fmt.Scan(&description)
			fmt.Print("Enter due date (YYYY-MM-DD): ")
			fmt.Scan(&dueDateStr)
			fmt.Print("Enter reminder (next day/week/month): ")
			fmt.Scan(&reminder)
			addChore(&manager, description, dueDateStr, reminder)
		case 2:
			listChores(manager)
		case 3:
			var id int
			fmt.Print("Enter chore ID to delete: ")
			fmt.Scan(&id)
			deleteChore(&manager, id)
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
