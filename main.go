package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// checkDistribution()
	pos := 0
	score := 0
	renderGameboard(0)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Roll Dice? [Y/N]")
		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToUpper(text)) != "Y" {
			fmt.Println("text: ", text)
			fmt.Println("Thanks for playing!")
			fmt.Println("Your final score is: ", score)
			return
		}
		fmt.Println(text)
		diceValue := rollDice(2)
		pos = (pos + diceValue) % 26
		score += 125

		switch pos {
		case pos:
			renderGameboard(pos)
		}
		fmt.Println("pos: ", pos)
		fmt.Printf("you rolled: %v\n", diceValue)
	}
}

func rollDice(number int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	dice1 := r1.Intn(6) + 1

	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	dice2 := r2.Intn(6) + 1
	return dice1 + dice2
}

func checkDistribution() {
	type diceValues struct {
		two    int
		three  int
		four   int
		five   int
		six    int
		seven  int
		eight  int
		nine   int
		ten    int
		eleven int
		twelve int
	}

	newDiceValues := diceValues{}

	for i := 0; i < 500; i++ {
		diceValue := rollDice(2)
		switch diceValue {
		case 2:
			newDiceValues.two++
		case 3:
			newDiceValues.three++
		case 4:
			newDiceValues.four++
		case 5:
			newDiceValues.five++
		case 6:
			newDiceValues.six++
		case 7:
			newDiceValues.seven++
		case 8:
			newDiceValues.eight++
		case 9:
			newDiceValues.nine++
		case 10:
			newDiceValues.ten++
		case 11:
			newDiceValues.eleven++
		case 12:
			newDiceValues.twelve++
		}
	}
	fmt.Println(strings.Repeat("*", newDiceValues.two))
	fmt.Println(strings.Repeat("*", newDiceValues.three))
	fmt.Println(strings.Repeat("*", newDiceValues.four))
	fmt.Println(strings.Repeat("*", newDiceValues.five))
	fmt.Println(strings.Repeat("*", newDiceValues.six))
	fmt.Println(strings.Repeat("*", newDiceValues.seven))
	fmt.Println(strings.Repeat("*", newDiceValues.eight))
	fmt.Println(strings.Repeat("*", newDiceValues.nine))
	fmt.Println(strings.Repeat("*", newDiceValues.ten))
	fmt.Println(strings.Repeat("*", newDiceValues.eleven))
	fmt.Println(strings.Repeat("*", newDiceValues.twelve))
}

func renderGameboard(value int) {
	switch value {
	case 0:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|  X  |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 1:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|  X  |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 2:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|  X  |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 3:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|  X  |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 4:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|  X  |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 5:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |  X  |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 6:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |  X  |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 7:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |  X  |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 8:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |  X  |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 9:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |  X  |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 10:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |  X  |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 11:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |  X  |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 12:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |  X  |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 13:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |  X  |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 14:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |  X  |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 15:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |  X  |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 16:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |  X  |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 17:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |     |  X  |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 18:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |     |  X  |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 19:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |     |  X  |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 20:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |     |  X  |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 21:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |     |  X  |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 22:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |     |  X  |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 23:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |     |  X  |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 24:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |     |  X  |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	case 25:
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |     |     |     |     |     |     |     |     |     |")
		fmt.Println("J--J--J-----|-----|-----S--S--S-----|-----|-----|-----H--H--H")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("|-----|                                               |-----|")
		fmt.Println("|     |                                               |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
		fmt.Println("|     |  X  |     |     |     |     |     |     |     |     |")
		fmt.Println("$--$--$-----|-----|-----S--S--S-----|-----|-----|-----|-----|")
	}

}
