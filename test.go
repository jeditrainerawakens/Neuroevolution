package main

import (
	"fmt"

	"github.com/CRRDerek/Neuroevolution/evolution"
	"github.com/CRRDerek/Neuroevolution/games"
	"github.com/CRRDerek/Neuroevolution/neuralnetwork"
)

func main() {
	//testXOR()
	//testTicTacToe()
	testCheckers()
}

// Seed a population of networks capable of learning XOR and then run neuroevolution
// on the XOR game.
func testXOR() {
	// Seed the initial population
	pop := make([]games.Agent, 100)
	for i := 0; i < 100; i++ {
		pop[i] = neuralnetwork.RandomNetwork(3, 4, 1)
	}

	evolution.EvolveAgents(games.XorGame, games.XorGamePlayerMaker,
		2000, 256, pop)

}

// Seed a population of networks capable of learning Tic Tac Toe (input size 28,
// output size 9) and run neuroevolution to produce an agent that has evolved to
// play tic tac toe.
//
// Run tic tac toe games against the user indefinitely once the evolved agent is ready.
func testTicTacToe() {
	// Seed the initial population
	pop_size := 256
	pop := make([]games.Agent, pop_size)
	for i := 0; i < pop_size; i++ {
		pop[i] = neuralnetwork.RandomNetwork(28, 56, 9)
	}

	// Evolve an agent capable of playing
	evolved_agent := evolution.EvolveAgents(games.TicTacToe, games.TicTacToePlayerMaker,
		10000, 1024, pop)

	fmt.Println("Training complete!")

	// Play tic tac toe against the user
	for {
		victor := games.TicTacToe(games.TicTacToePlayerMaker(evolved_agent), games.HumanTicTacToePlayer)
		if victor == -1 {
			fmt.Println("\n\nYou win!")
		} else if victor == 0 {
			fmt.Println("\n\nDraw!")
		} else if victor == 1 {
			fmt.Println("\n\nYou lose!")
		}
	}

}

// Seed a population of networks capable of learning Checkers (input size 65,
// output size 24) and run neuroevolution to produce an agent that has evolved to
// play Checkers.
//
// Run Checkers games against the user indefinitely once the evolved agent is ready.
func testCheckers() {
	// Seed the initial population
	pop_size := 100
	pop := make([]games.Agent, pop_size)
	for i := 0; i < pop_size; i++ {
		pop[i] = neuralnetwork.RandomNetwork(65, 130, 24)
	}

	// Run neuroevolution to produce an agent. The checkers games used by the
	// evolutionary algorithm will be cut off after 100 moves to prevent
	// random players from prolonging the game indefinitely.
	evolved_agent := evolution.EvolveAgents(games.MakeCheckers(100), games.CheckersPlayerMaker,
		512, 64, pop) // Each member of the population will be tested at maximum 64 times.
	// After 512 generations the algorithm concludes if it hasn't already spawned
	// an agent that can win 64 times for 10 generations.
	fmt.Println("Training complete!")

	// Play checkers against the user indefinitely
	for {
		victor := games.Checkers(games.CheckersPlayerMaker(evolved_agent), games.HumanCheckersPlayer)
		if victor == -1 {
			fmt.Println("\n\nYou win!")
		} else if victor == 0 {
			fmt.Println("\n\nDraw!")
		} else if victor == 1 {
			fmt.Println("\n\nYou lose!")
		}
	}

}

// Test the user interface of checkers against a random player
func testRandomCheckers() {
	for {
		victor := games.Checkers(games.HumanCheckersPlayer, games.RandomPlayer)
		if victor == 1 {
			fmt.Println("\n\nYou win!")
		} else if victor == 0 {
			fmt.Println("\n\nDraw!")
		} else if victor == -1 {
			fmt.Println("\n\nYou lose!")
		}
	}
}
