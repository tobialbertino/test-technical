package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Player struct {
	Dice      []int
	TotalDice int
	Score     int
	Name      int
}

type Game struct {
	Players    []Player
	Round      int
	PlayerDone []Player
}

// InitiateGame, initialize the number of players and totaldice.
func (game *Game) InitiateGame(totalPlayer, totalDice int) {
	fmt.Printf("Pemain = %v, Dadu = %v\n", totalPlayer, totalDice)
	println("==================")

	// loop through totalPlayer, add it
	for i := 0; i < totalPlayer; i++ {
		game.Players = append(game.Players, Player{
			TotalDice: totalDice,
			Name:      i,
		})
	}

	// loops through game players
	for i := 0; i < len(game.Players); i++ {

		// loop append slice, then initialize totalDice by 0
		for j := 0; j < game.Players[i].TotalDice; j++ {
			game.Players[i].Dice = append(game.Players[i].Dice, 0)
		}
	}

}

// Roll, shuffle each player's dice
func (game *Game) Roll(number int) {
	fmt.Printf("Giliran %v lempar dadu:\n", number)

	// loops through game players
	for i := 0; i < len(game.Players); i++ {
		fmt.Printf("     Pemain #%v (%v): ", i, game.Players[i].Score)

		// Loop Generate random Dice number range (0, 6)
		for j := 0; j < len(game.Players[i].Dice); j++ {
			game.Players[i].Dice[j] = rand.Intn(6) + 1
			fmt.Printf("%v, ", game.Players[i].Dice[j])
		}
		println()
	}
}

// delelete an element from a slice.
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func removePlayer(slice []Player, s int) []Player {
	return append(slice[:s], slice[s+1:]...)
}

// Evaluate1, add score where dice = 6.
func (game *Game) Evaluate_1() {

	// loops through game players
	for i := 0; i < len(game.Players); i++ {

		// loop throgh dice player, count point score
		for j := 0; j < len(game.Players[i].Dice); j++ {

			// if dice == 6, add score, remove dice, resize totalDice
			if game.Players[i].Dice[j] == 6 {
				game.Players[i].Score += 1
				game.Players[i].Dice = remove(game.Players[i].Dice, j)
				game.Players[i].TotalDice -= 1
			}
		}
	}
}

// FIX ME, ada proses yang kedouble, update: looks good
// Evaluate_2 pass dice to its side when value = 1,
func (game *Game) Evaluate_2() {
	// copyGame, get fixed data before, len of players and len of player dice's

	for i := 0; i < len(game.Players); i++ {

		// loop throgh dice player, count point score
		for j := 0; j < game.Players[i].TotalDice; j++ {

			// if dice == 1, pass the dice to its side
			if game.Players[i].Dice[j] == 1 {
				game.Players[i].Dice = remove(game.Players[i].Dice, j)
				game.Players[i].TotalDice -= 1

				// if last player, pass the dice to its side from the beginning
				if len(game.Players)-1 == i {
					game.Players[0].Dice = append(game.Players[0].Dice, 1)
				} else {
					// pass the dice to its side
					game.Players[i+1].Dice = append(game.Players[i+1].Dice, 1)
				}
			}
		}

	}

	// update total dice
	for i := 0; i < len(game.Players); i++ {
		// loop throgh dice player, count point score
		for j := 0; j < len(game.Players[i].Dice); j++ {
			game.Players[i].TotalDice = len(game.Players[i].Dice)
		}
	}
}

func (game *Game) PrintEvaluate() {
	println("Setelah evaluasi:")

	// loops through game players
	for i := 0; i < len(game.Players); i++ {

		// loop throgh dice player, count point score
		fmt.Printf("     Pemain #%v (%v): ", i, game.Players[i].Score)
		sort.Ints(game.Players[i].Dice)

		// after score evaluated
		// loop throgh dice player
		for j := 0; j < len(game.Players[i].Dice); j++ {
			fmt.Printf("%v, ", game.Players[i].Dice[j])
		}
		println()
	}
}

// CheckDice, if player is end, save score
func (game *Game) CheckDice() {
	for i := 0; i < len(game.Players); i++ {
		if len(game.Players[i].Dice) == 0 {
			game.PlayerDone = append(game.PlayerDone, game.Players[i])
			game.Players = removePlayer(game.Players, i)
		}
	}
}

func (game *Game) PlayLoop() {
	for {
		game.Roll(game.Round)
		game.Evaluate_1()
		game.Evaluate_2()
		game.PrintEvaluate()
		game.CheckDice()
		if len(game.Players) == 1 {
			break
		}
		game.Round++
	}
	max := 0
	name := 0
	game.PlayerDone = append(game.PlayerDone, game.Players...)
	for _, v := range game.PlayerDone {
		if max < v.Score {
			max = v.Score
			name = v.Name
		}
	}
	fmt.Println("Finish")
	fmt.Printf("The best score is Player %v with score %v", name, max)

}

func main() {
	game := Game{}
	game.InitiateGame(4, 5)
	game.PlayLoop()
}
