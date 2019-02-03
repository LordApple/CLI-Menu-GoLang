package main

import (
	"bufio"
	"os"
)

func ask(array []string) string {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inp := scanner.Text()
		for _, item := range array {
			if inp == item {
				return inp
			}
		}

	}

}
