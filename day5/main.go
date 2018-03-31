package main

import (
	"fmt"
	"strings"
)

func main() {
	var coins int = 50
	var users []string = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	result := allocCoins(users, coins)
	fmt.Println(result)

}

func allocCoins(users []string, conis int) map[string]int {
	var result = make(map[string]int, len(users))
	for _, name := range users {
		if ((strings.IndexByte(name, 'a')) != -1) || ((strings.IndexByte(name, 'A')) != -1) {
			result[name]++
		}
		if ((strings.IndexByte(name, 'e')) != -1) || ((strings.IndexByte(name, 'E')) != -1) {
			result[name]++
		}
		if ((strings.IndexByte(name, 'i')) != -1) || ((strings.IndexByte(name, 'I')) != -1) {
			result[name] += 2
		}
		if ((strings.IndexByte(name, 'o')) != -1) || ((strings.IndexByte(name, 'O')) != -1) {
			result[name] += 3
		}
		if ((strings.IndexByte(name, 'u')) != -1) || ((strings.IndexByte(name, 'U')) != -1) {
			result[name] += 5
		}
	}
	return result
}
