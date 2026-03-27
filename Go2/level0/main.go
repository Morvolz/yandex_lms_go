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
	for _, line := range lines {
		if strings.HasPrefix(line, "TICKET") {
		} else {
			continue
		}
		var mTicket Ticket
		mTicket.Ticket = line[:strings.Index(line, "_")]
		line = line[strings.Index(line, "_")+1:]

		mTicket.User = line[:strings.Index(line, "_")]
		line = line[strings.Index(line, "_")+1:]

		mTicket.Status = line[:strings.Index(line, "_")]
		line = line[strings.Index(line, "_")+1:]

		parsedTime, err := time.Parse("2006-01-02", line)
		if err != nil {
			continue
		}

		mTicket.Date = parsedTime
		if user != nil {
			if *user != mTicket.User {
				// fmt.Print(mTicket.User, "  ", *user, "\n")
				continue
			}
		}
		if status != nil {
			if mTicket.Status != *status {
				continue
			}
			// fmt.Print(mTicket.Status, "\n")
		}

		// fmt.Printf("%d: %q %q %q\n", i, mTicket.Ticket, mTicket.User, mTicket.Status)
		tickets = append(tickets, mTicket)
	}
	return tickets
}

func main() {
	chatHistory := "TICKET-12345_Паша Попов_Готово_2024-01-01\nTICKET-12346_Иван Иванов_В работе_2024-01-02\nTICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03\nTICKET-12348_Паша Попов_В работе_2024-01-04"
	user := "Паша Попов"
	tasks := GetTasks(chatHistory, &user, nil)
	fmt.Println(tasks)

	status := "В работе"
	tasks = GetTasks(chatHistory, nil, &status)
	fmt.Println(tasks)
}
