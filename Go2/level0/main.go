package main

import (
	"fmt"
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

func GetTasks(text string, user *string, status *string) []Ticket {
	var tickets []Ticket
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		fmt.Printf("%d: %q\n", i, line)
	}
	return tickets
}

func main() {
	chatHistory := "TICKET-12345_Паша Попов_Готово_2024-01-01\nTICKET-12346_Иван Иванов_В работе_2024-01-02\nTICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03\nTICKET-12348_Паша Попов_В работе_2024-01-04"
	user := "Паша Попов"
	tasks := GetTasks(chatHistory, &user, nil)
	fmt.Print(tasks)
}
