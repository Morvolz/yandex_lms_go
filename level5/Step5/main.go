package main

import (
	"fmt"
	"slices"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func (p *Player) calculateRating() {
	if p.Misses == 0 {
		p.Rating = float64(p.Goals) + float64(p.Assists)/2.
	} else {
		p.Rating = (float64(p.Goals) + float64(p.Assists)/2.) / float64(p.Misses)
	}
}

func NewPlayer(name string, goals, misses, assists int) Player {
	p := Player{
		Name:    name,
		Goals:   goals,
		Misses:  misses,
		Assists: assists,
	}
	p.calculateRating()
	return p
}

func goalsSort(players []Player) []Player {

	slices.SortFunc(players, func(a, b Player) int {
		switch {
		case a.Goals < b.Goals:
			return 1
		case a.Goals > b.Goals:
			return -1
		case a.Name > b.Name:
			return 1
		case a.Name < b.Name:
			return -1
		default:
			return 0
		}
	})
	return players
}

func ratingSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		switch {
		case a.Rating < b.Rating:
			return 1
		case a.Rating > b.Rating:
			return -1
		case a.Name > b.Name:
			return 1
		case a.Name < b.Name:
			return -1
		default:
			return 0
		}
	})
	return players
}

func gmSort(players []Player) []Player {
	slices.SortFunc(players, func(a, b Player) int {
		var gma, gmb float64
		gma = float64(a.Goals) / float64(a.Misses+1)
		gmb = float64(b.Goals) / float64(b.Misses+1)
		switch {
		case gma < gmb:
			return 1
		case gma > gmb:
			return -1
		case a.Name > b.Name:
			return 1
		case a.Name < b.Name:
			return -1
		default:
			return 0
		}
	})
	return players
}

func main() {
	// Создаём нескольких игроков для тестирования
	players := []Player{
		NewPlayer("Арсений", 10, 5, 7),  // рейтинг: (10+7/2)/5 = (10+3.5)/5 = 13.5/5=2.7
		NewPlayer("Дмитрий", 8, 3, 4),   // (8+2)/3 = 10/3 ≈ 3.333
		NewPlayer("Алексей", 12, 0, 2),  // промахов нет: 12+1 = 13
		NewPlayer("Борис", 12, 2, 5),    // (12+2.5)/2 = 14.5/2=7.25
		NewPlayer("Владимир", 7, 7, 1),  // (7+0.5)/7 = 7.5/7 ≈ 1.0714
		NewPlayer("Григорий", 10, 5, 2), // (10+1)/5 = 11/5=2.2
		NewPlayer("Александр", 8, 4, 6), // (8+3)/4 = 11/4=2.75
	}

	// Выводим исходный список с рейтингом
	fmt.Println("Исходный список игроков:")
	printPlayers(players)

	// Сортировка по голам
	sortedByGoals := goalsSort(players)
	fmt.Println("\nСортировка по убыванию голов (при равенстве — по имени):")
	printPlayers(sortedByGoals)

	// Сортировка по рейтингу
	sortedByRating := ratingSort(players)
	fmt.Println("\nСортировка по убыванию рейтинга (при равенстве — по имени):")
	printPlayers(sortedByRating)

	// Сортировка по отношению голов к промахам
	sortedByGM := gmSort(players)
	fmt.Println("\nСортировка по убыванию отношения голов к промахам (при равенстве — по имени):")
	printPlayers(sortedByGM)
}

// printPlayers выводит список игроков в читаемом формате
func printPlayers(players []Player) {
	fmt.Printf("%-12s %5s %5s %5s %8s\n", "Имя", "Голы", "Промахи", "Помощь", "Рейтинг")
	fmt.Println("--------------------------------------------------")
	for _, p := range players {
		fmt.Printf("%-12s %5d %5d %5d %8.3f\n", p.Name, p.Goals, p.Misses, p.Assists, p.Rating)
	}
}
