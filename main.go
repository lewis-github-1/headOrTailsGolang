package main

import (
	"fmt"
	"math/rand"
	"time"
)

func flipCoin() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return (random.Intn(2) + 1)
}

func userInput() (int, error) {
	var choice int
	fmt.Println("Press 1 for heads and 2 for tails.")
	_, err := fmt.Scan(&choice)
	if err != nil {
		return 0, err
	}
	return choice, err
}

func numberToHeadsOrTails(choiceInt, resultInt int) (string, string, error) {
	var choiceString, resultString string

	switch choiceInt {
	case 1:
		choiceString = "heads"
	case 2:
		choiceString = "tails"
	default:
		return "", "", fmt.Errorf("invalid choice: %d", choiceInt)
	}

	switch resultInt {
	case 1:
		resultString = "heads"
	case 2:
		resultString = "tails"
	default:
		return "", "", fmt.Errorf("invalid result: %d", resultInt)
	}

	return choiceString, resultString, nil
}

func main() {
	counter := 0
	counterWin := 0
	counterTotal := 0
	for {
		for {
			userChoice, err := userInput()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			result := flipCoin()

			choiceString, resultString, err := numberToHeadsOrTails(userChoice, result)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println("You chose", choiceString)
			fmt.Println("The coin landed on", resultString)

			if userChoice == result {
				fmt.Println("You chose correctly")
				counterWin++
			} else {
				fmt.Println("You chose poorly.")
			}
			counter++
			counterTotal++
			fmt.Printf("You have won %.2f%% of games. %d / %d \n", (float64(counterWin) / float64(counterTotal) * 100), counterWin, counterTotal)
			fmt.Println()

			if counter >= 10 {
				counter = 0
				break
			}

		}
		fmt.Println("Game over!")
		fmt.Println("Keep playing? (y/n) ")
		var playAgain string
		fmt.Scan(&playAgain)
		if playAgain != "y" {
			fmt.Println("Are you sure you want to quit? (y/n)")
			var confirmQuit string
			fmt.Scan(&confirmQuit)
			if confirmQuit == "y" {
				fmt.Println("Thank you for playing!")
				break
			} else {
				fmt.Println("Continuing the game...")
				continue
			}
		}
	}
}

// BUG: count is getting off of multiples of 10 after y/n. continue or break in wrong place?
// Make a menu

// Add to website
// Make a leaderboard
// Add images
// Multiplayer?
// Chat room?
