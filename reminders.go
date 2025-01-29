package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func sendEmailReminder(chore Chore, email string) {
	from := "your_email@example.com"
	password := "your_email_password"
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	message := fmt.Sprintf("Subject: Chore Reminder\n\nDon't forget to complete the chore: %s\nDue Date: %s", chore.Description, chore.DueDate.Format("2006-01-02"))

	auth := smtp.PlainAuth("", from, password, smtpHost)
	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, []byte(message)); err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}
	fmt.Println("Reminder email sent successfully!")
}
