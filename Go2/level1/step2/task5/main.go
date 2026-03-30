package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	f, err := os.Open(inputFileName)
	fileScanner := bufio.NewScanner(f)
	var logs []string
	for fileScanner.Scan() {
		s := fileScanner.Text()
		parts := strings.SplitN(s, " ", 2)
		dateStr := parts[0]
		date, _ := time.Parse("02.01.2006", dateStr)
		if (date.After(start) || date.Equal(start)) && (date.Before(end) || date.Equal(end)) {
			logs = append(logs, s)
		}
	}
	if len(logs) == 0 {
		return nil, errors.New("Pu-pu")
	}
	return logs, err
}
